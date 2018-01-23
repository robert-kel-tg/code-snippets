package channels

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

const MAX = 1000

// Example of two goroutines working concurently
func DemoChannels0() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		work := make(chan int, MAX)
		result := make(chan int)

		// Pass directional; receiving channel
		go func(work chan<- int) {
			for i := 1; i < MAX; i++ {
				if (i%3) == 0 || (i%5) == 0 {
					work <- i
				}
			}
			close(work)
		}(work)

		// Pass directional; receiving channel
		go func(result chan<- int) {
			r := 0
			for i := range work {
				r = r + i
			}
			result <- r
		}(result)

		w.Write([]byte(fmt.Sprintf("%v", <-result)))
	}
}

// Streaming data pattern
// A natural use of channels is to stream data from one goroutine to another
// https://www.safaribooksonline.com/library/view/go-design-patterns/9781788390552/ch09s03.html
func DemoChannels1() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		data := []string{
			"The yellow fish swims slowly in the water",
			"The brown dog barks loudly after a drink ...",
			"The dark bird bird of prey lands on a small ...",
		}

		histogram := make(map[string]int)

		words := words(data)

		//Pulls the data from the channel
		//Checks the open status of the channel
		//If closed, break out of the loop
		//Otherwise record histogram
		for {
			word, opened := <-words
			if !opened {
				break
			}
			histogram[word]++
		}

		for k, v := range histogram {
			w.Write([]byte(fmt.Sprintf("%s\t %d\n", k, v)))
		}
	}
}

// returns receive-only channel
// The consumer function, in this instance main(), receives the data emitted by the generator function
func words(data []string) <-chan string {
	out := make(chan string)

	go func() {
		defer close(out)

		for _, line := range data {
			words := strings.Split(line, " ")
			for _, word := range words {
				word = strings.ToLower(word)

				out <- word
			}
		}
	}()

	return out
}

func DemoSelect() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		helloChannel := make(chan string, 1)
		byeChannel := make(chan string, 1)
		quitChannel := make(chan bool)

		go receiver(helloChannel, byeChannel, quitChannel)

		go sendString(helloChannel, "Hey there!")
		time.Sleep(time.Second)
		go sendString(byeChannel, "Bye!")

		<-quitChannel
	}
}

func sendString(sendChannel chan<- string, data string) {
	sendChannel <- data
}

func receiver(helloChannel <-chan string, byeChannel <-chan string, quitChannel chan<- bool) {
	for {
		select {
		case msg := <-helloChannel:
			println(msg)
		case msg := <-byeChannel:
			println(msg)
		case <-time.After(time.Second * 2):
			println("Nothing received in 2 seconds; Exiting!")
			quitChannel <- true
			break
		}
	}
}
