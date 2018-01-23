package usertype

import (
	"fmt"
	"math/rand"
	"net/http"
	"sync"

	"github.com/robertke/orders-service/pkg/decorator"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
)

var wg sync.WaitGroup

type notifier interface {
	notify()
}

type User struct {
	ID       uuid.UUID
	Name     string
	Username string
}

// u User with * as it's pointer receiver
func (u *User) changeNameTo(name string) {
	u.Name = name
}

// Run example
// runtime.GOMAXPROCS(runtime.NumCPU())
// on main.go add usertype.Run()
func Run() {
	crt := make(chan string)
	wg.Add(3)
	go check(crt, "Aaa")
	go check(crt, "Bbb")
	go check(crt, "Ccc")

	tests := []User{}
	tests = append(tests, User{
		ID:       uuid.NewV1(),
		Username: "maas",
		Name:     "test1",
	})
	tests = append(tests, User{
		ID:       uuid.NewV1(),
		Username: "wews",
		Name:     "test2",
	})

	for _, test := range tests {
		crt <- test.Name
	}

	wg.Wait()
}

func UserTypeHandler(w http.ResponseWriter, r *http.Request) {
	user := User{
		ID:       uuid.NewV1(),
		Name:     "Test",
		Username: "Usrname",
	}

	user.changeNameTo("NewName")
	sendNotification(user)

	if err := decorator.Exec(); err != nil {
		log.WithError(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write([]byte(fmt.Sprintf("UserType request")))
}

//https://www.safaribooksonline.com/library/view/go-in-action/9781617291784/kindle_split_013.html
// u User without * as it's value receiver
func (u User) notify() {
	fmt.Printf("User %s<%s>\n", u.Name, u.Username)
}

func sendNotification(n notifier) {
	n.notify()
}

func check(all chan string, kind string) {

	defer wg.Done()

	msgs, ok := <-all

	if !ok {
		fmt.Printf("Closed %s\n", kind)
		return
	}

	n := rand.Intn(100)
	if n%13 == 0 {
		fmt.Printf("Player %s Missed\n", kind)

		// Close the channel to signal we lost.
		close(all)
		return
	}

	fmt.Printf("%s does %s\n", kind, msgs)

	all <- msgs
}
