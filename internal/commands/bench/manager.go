package bench

import (
	"github.com/amirhnajafiz/cemq/internal/mqtt"
	"github.com/amirhnajafiz/cemq/pkg/model"
)

// Manager handles benchmakr commands
type Manager interface {
	Benchmark(bench *model.Benchmark) *model.Upshot
}

func New(conn mqtt.CLI) Manager {
	return &bench{
		conn: conn,
	}
}
