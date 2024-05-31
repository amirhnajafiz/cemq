package load

import (
	"fmt"
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

	return nil
}

func (l *load) handleTopic(topic string, pubs int, subs int, period time.Duration) error {
	return nil
}
