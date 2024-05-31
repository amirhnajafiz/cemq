package model

// Load stores the input data to generate a load on an EMQX cluster
type Load struct {
	PublishIntervals int
	Publishers       int
	Subscribers      int
	Topics           int
}
