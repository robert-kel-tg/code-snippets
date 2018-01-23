# Runner

```go
const timeout = 3 * time.Second

log.Println("Starting work")

		runner := concurency.NewRunner(timeout)
		runner.Add(createTask(), createTask(), createTask())

		if err := runner.Start(); err != nil {
			switch err {
			case concurency.ErrTimeout:
				log.Println("Terminating due to timeout")
				os.Exit(1)
			case concurency.ErrInterrupt:
				log.Println("Terminating due to interrupt")
				os.Exit(2)
			}

			log.Println("Process ended")
		}

// Returns example task that sleeps specified number of time
func createTask() func(int) {
	return func(id int) {
		log.Printf("Processor - Task #%d.", id)
		time.Sleep(time.Duration(id) * time.Second)
	}
}
```

# Pool

```go
const (
	maxGoroutines   = 25
	pooledResources = 2
)

type dbConnection struct {
	ID int32
}

func (dbConn *dbConnection) Close() error {
	log.Println("Close: Connection", dbConn.ID)

	return nil
}

var idCounter int32

func createConnection() (io.Closer, error) {

	id := atomic.AddInt32(&idCounter, 1)
	log.Println("Create: New Connection", id)
	return &dbConnection{id}, nil
}



var wg sync.WaitGroup
	wg.Add(maxGoroutines)

	pool, err := concurency.New(createConnection, pooledResources)
	if err != nil {
		log.WithError(err).Error("Coudn't connect to the resource")
	}

	for query := 0; query < maxGoroutines; query++ {
		go func(q int) {
			performQueries(q, pool)
			wg.Done()
		}(query)
	}

	wg.Wait()

	log.Println("Shutdown Program")
	pool.Close()


func performQueries(query int, p *concurency.Pool) {
	conn, err := p.Retrieve()
	if err != nil {
		log.Println(err)
		return
	}

	defer p.Release(conn)

	// Simulate query response time
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	log.Printf("QID[%d] CID[%d]\n", query, conn.(*dbConnection).ID)
}
```


# Work

```go
var names = []string{
	"steve",
	"bob",
	"mary",
	"therese",
	"jason",
}

type namePrinter struct {
	name string
}

func (m *namePrinter) Task() {
	log.Println(m.name)
	time.Sleep(time.Second)
}

p := concurency.NewPool(2)

var wg sync.WaitGroup
wg.Add(100 * len(names))

for i := 0; i < 100; i++ {
	for _, name := range names {
		np := namePrinter{
			name: name,
		}

		go func() {
			p.Run(&np)
			wg.Done()
		}()
	}
}
wg.Wait()
p.Shutdown()
```