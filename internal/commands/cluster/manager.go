package cluster

import "github.com/amirhnajafiz/cemq/internal/mqtt"

// Manager handles the cluster availability status
type Manager interface {
	CheckHealth() string
	CheckConnection() string
}

func New(conn mqtt.CLI) Manager {
	return &cluster{
		conn: conn,
	}
}
