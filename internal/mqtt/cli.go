package mqtt

import (
	"log"
	"os"

	"github.com/amirhnajafiz/cemq/pkg/model"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// CLI is used to simplfy the usage of MQTT cluster by CEMQ components
type CLI interface {
	// Connect is used to establish a new connection before running crutial methods like Topics, Publish, ...
	Connect() error
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

// NewCLI creates a cli instance with client options in order to connect to the current EMQX cluster;
// the config instance will be read from config file and passed as a model.config type
func NewCLI(cfg *model.Config, debug bool) CLI {
	if debug {
		mqtt.DEBUG = log.New(os.Stdout, "", 0)
		mqtt.ERROR = log.New(os.Stdout, "", 0)
	}

	opts := mqtt.NewClientOptions()
	opts.AddBroker(cfg.URL())

	if len(cfg.Username) > 0 {
		opts.SetUsername(cfg.Username)
	}
	if len(cfg.Password) > 0 {
		opts.SetPassword(cfg.Password)
	}

	client := mqtt.NewClient(opts)

	return &cli{
		conn: client,
	}
}
