package model

import "fmt"

type StatType string

const (
	PubStat StatType = "publish"
	SubStat StatType = "subscribe"
)

// Upshot stores a benchmark result
type Upshot struct {
	ID    string
	Topic string
	Stats []*Stat
}

func (u Upshot) ToString() string {
	return fmt.Sprintf("Benchmark %s: topic %s", u.ID, u.Topic)
}

// Stat is a single entity that holds benchmark data
type Stat struct {
	ID         string
	Type       string
	Throughput float64
	Mps        float64
}

func (s Stat) ToString() string {
	return fmt.Sprintf("[%s] %s: %f mps, %f kbyte/s", s.ID, s.Type, s.Mps, s.Throughput)
}
