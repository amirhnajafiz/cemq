package model

// Upshot stores a benchmark result
type Upshot struct {
	ID    string
	Topic string
	Pubs  []*PubStat
	Subs  []*SubStat
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
