package model

import (
	"time"
)

type Result struct {
	ID        string
	Duration  time.Time
	Benchmark *Benchmark
	Pubs      []*PubStat
	Subs      []*SubStat
}

type PubStat struct {
	ID              string
	WriteThroughput float64
	WriteMps        float64
}

type SubStat struct {
	ID             string
	ReadThroughput float64
	ReadMps        float64
}
