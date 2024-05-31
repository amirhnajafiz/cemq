package bench

import (
	"github.com/amirhnajafiz/cemq/internal/mqtt"
	"github.com/amirhnajafiz/cemq/pkg/model"
)

type load struct {
	conn mqtt.CLI
}

// Generate method creates publishers and subscribers according to the number of topics
func (l load) Generate(input *model.Load) error {
	return nil
}
