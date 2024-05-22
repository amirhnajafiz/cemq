package bench

import (
	"log"

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
				<-channel
				counter.Inc()
			}
		}()
	}

	counter.Wait()
}

// generatePublishters create a number of publishers to publish msgs over a topic
func (b bench) generatePublishters(topic string, msgs, number int) {
	counter := utils.NewSafeCounter(msgs)
	channel := make(chan bool)

	for i := 0; i < number; i++ {
		go func() {
			for {
				<-channel

				if err := b.conn.Publish(topic, []byte("testing message to send")); err != nil {
					log.Println(err)

					channel <- true
				} else {
					counter.Inc()
				}
			}
		}()
	}

	for i := 0; i < msgs; i++ {
		channel <- true
	}

	counter.Wait()
}
