package bench

import (
	"github.com/amirhnajafiz/cemq/internal/mqtt"
	"github.com/amirhnajafiz/cemq/pkg/model"
)

type bench struct {
	conn mqtt.CLI
}

// Benchmark method handles a benchmark request and returns an upshot
func (b bench) Benchmark(bench *model.Benchmark) *model.Upshot {
	return nil
}
