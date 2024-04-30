package model

import "time"

type Result struct {
	ID        string
	Duration  time.Time
	Benchmark *Benchmark
	Micros    []*MicroResult
}

type MicroResult struct {
	ID              string
	Type            string
	ReadThroughput  float64
	WriteThroughput float64
	ReadMps         float64
	WriteMps        float64
}
