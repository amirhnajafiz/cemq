package cluster

type Manager interface {
	CheckHealth() string
	CheckConnection() string
}
