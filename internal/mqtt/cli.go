package mqtt

type CLI interface {
	// emqx Cluster methods
	CheckConnection() (string, error)
	CheckHealth() (string, error)
	Topics() ([]string, error)
	// emqx Benchmakr methods
	Publish(topic, message string) error
	Subscribe(topic string) chan string
}
