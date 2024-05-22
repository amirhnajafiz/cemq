package bench

import (
	"github.com/amirhnajafiz/cemq/internal/mqtt"
	"github.com/amirhnajafiz/cemq/internal/utils"
	"github.com/amirhnajafiz/cemq/pkg/model"
)

type bench struct {
	conn mqtt.CLI
}

// Benchmark method handles a benchmark request and returns an upshot
func (b bench) Benchmark(bench *model.Benchmark) *model.Upshot {
	return nil
}

// generateConsumers create a number of consumers to get msgs from a topic
func (b bench) generateConsumers(topic string, msgs, number int) {
	counter := utils.NewSafeCounter(msgs)
}

// generatePublishters create a number of publishers to publish msgs over a topic
func (b bench) generatePublishters(topic string, msgs, number int) {

}
