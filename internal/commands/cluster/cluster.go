package cluster

import (
	"github.com/amirhnajafiz/cemq/internal/mqtt"
)

type cluster struct {
	conn mqtt.CLI
}

// CheckHealth measures the cluster status by a ping-pong mechanism
func (c cluster) CheckHealth() string {
	if err := c.conn.Connect(); err != nil {
		return err.Error()
	}

	msg, err := c.conn.CheckHealth()
	if err != nil {
		return err.Error()
	}

	return msg
}

// CheckConnection only checks the EMQX cluster availability
func (c cluster) CheckConnection() string {
	msg, err := c.conn.CheckConnection()
	if err != nil {
		return err.Error()
	}

	return msg
}
