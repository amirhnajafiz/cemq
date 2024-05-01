package bench

import "github.com/amirhnajafiz/cemq/pkg/model"

type Manager interface {
	Benchmark(bench *model.Benchmark) *model.Result
}
