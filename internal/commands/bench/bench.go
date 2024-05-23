package bench

import (
	"log"
	"sync"
	"time"

	"github.com/amirhnajafiz/cemq/internal/mqtt"
	"github.com/amirhnajafiz/cemq/internal/utils"
	"github.com/amirhnajafiz/cemq/pkg/model"
)

type bench struct {
	conn mqtt.CLI
}

// Benchmark method handles a benchmark request and returns an upshot
func (b bench) Benchmark(bench *model.Benchmark) *model.Upshot {
	go b.generateConsumers(bench.Topic, bench.Messages, bench.Subs)
	go b.generatePublishters(bench.Topic, bench.Messages, bench.Pubs)

	return nil
}

// generateConsumers create a number of consumers to get msgs from a topic
func (b bench) generateConsumers(topic string, msgs, number int) {
	counter := utils.NewSafeCounter(msgs)

	for i := 0; i < number; i++ {
		go func() {
			channel, err := b.conn.Subscribe(topic)
			if err != nil {
				log.Println(err)

				return
			}

			for {
				msg := model.CreateMessageFromBytes(<-channel)
				msg.TTL = int(time.Since(msg.Sent))

				counter.Inc()
			}
		}()
	}

	counter.Wait()
}

// generatePublishters create a number of publishers to publish msgs over a topic.
// creating maximum `number` go-routines
func (b bench) generatePublishters(topic string, msgs, number int) {
	// create a wg from the number of publishers
	var wg sync.WaitGroup

	// create a go-routine for each publisher to publish N messages
	for i := 0; i < number; i++ {
		wg.Add(1)

		go func() {
			index := 0
			for {
				// check for limit
				if index == msgs {
					break
				}

				// publish a message over the topic
				if err := b.conn.Publish(topic, model.Message{
					Sent: time.Now(),
				}.ToBytes()); err != nil {
					log.Println(err)
				} else {
					index++ // increase index if succeed
				}
			}

			// release wg
			wg.Done()
		}()
	}

	// wait for all go-routines
	wg.Wait()
}
