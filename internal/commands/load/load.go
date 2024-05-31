package load

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/amirhnajafiz/cemq/internal/mqtt"
	"github.com/amirhnajafiz/cemq/pkg/model"
)

type load struct {
	conn   mqtt.CLI
	lock   sync.Mutex
	errors []error
}

// Generate method creates publishers and subscribers according to the number of topics
func (l *load) Generate(input *model.Load) error {
	// create a waitgroup to manage sub-processes
	var wg sync.WaitGroup

	// create interval
	interval := time.Duration(input.PublishIntervals) * time.Millisecond

	for i := 1; i <= input.Topics; i++ {
		wg.Add(1)

		go func(index int) {
			topic := fmt.Sprintf("cemqx.test.%d", index)
			if err := l.handleTopic(topic, input.Publishers, input.Subscribers, interval); err != nil {
				l.lock.Lock()
				l.errors = append(l.errors, err)
				l.lock.Unlock()
			}

			wg.Done()
		}(i)
	}

	// wait for all topics to stop
	wg.Wait()

	// check for errors
	if len(l.errors) > 0 {
		return l.errors[0]
	}

	return nil
}

// handleTopic is used to create publishers and subscribers over a topic
func (l *load) handleTopic(topic string, pubs int, subs int, period time.Duration) error {
	var (
		killSwitch = make(chan int) // create a kill switch channel to sync workers on errors
		errorlog   error
		lock       sync.Mutex
	)

	// create subscribers
	for i := 0; i < subs; i++ {
		go func() {
			// subscribe to the topic
			channel, err := l.conn.Subscribe(topic)
			if err != nil {
				lock.Lock()
				errorlog = err
				lock.Unlock()

				killSwitch <- 1
			}

			// start getting data
			for {
				data := <-channel
				log.Printf("consume %d bytes\n", len(data))
			}
		}()
	}

	// create publishers
	for i := 0; i < pubs; i++ {
		go func() {
			// create a new ticker
			ticker := time.NewTicker(period)
			for range ticker.C {
				// publish a message over the topic
				if err := l.conn.Publish(topic, []byte("a message to be sent")); err != nil {
					lock.Lock()
					errorlog = err
					lock.Unlock()

					killSwitch <- 1
				}
			}
		}()
	}

	// wait for an error
	<-killSwitch

	return errorlog
}
