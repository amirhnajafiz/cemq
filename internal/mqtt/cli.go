package mqtt

import "github.com/amirhnajafiz/cemq/pkg/model"

// CLI is used to simplfy the usage of MQTT cluster by CEMQ components
type CLI interface {
	// emqx Cluster methods
	// CheckConnection is used to check the connection health to the current EMQX cluster
	CheckConnection() (string, error)
	// CheckHealth publishes over a testing topic in order to check the current EMQX cluster health
	CheckHealth() (string, error)
	// Topics aims to return the current topics that are available on the current EMQX cluster
	Topics() ([]string, error)
	// emqx Benchmakr methods
	// Publish is used to send a bytes data to the current EMQX cluster over a topic
	Publish(topic string, message []byte) error
	// Subscribe opens a channel in order to recieve events from a topic over the current EMQX cluster
	Subscribe(topic string) chan []byte
}

func NewCLI(cfg *model.Config) (CLI, error) {
	return &cli{}, nil
}
