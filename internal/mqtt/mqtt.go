package mqtt

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// cli is the core of mqtt package, in which we implemented the logic of methods that
// are used to communicate with a EMQX cluster
type cli struct {
	conn mqtt.Client
}

// connect is an internal method used to establish a connection to EMQX cluster
func (c cli) Connect() error {
	token := c.conn.Connect()
	defer func() {
		token.Done()
	}()

	if token.Wait() && token.Error() != nil {
		return token.Error()
	}

	return nil
}

func (c cli) CheckConnection() (string, error) {
	if err := c.Connect(); err != nil {
		return ErrConnection.Error(), err
	}

	return "connected!", nil
}

func (c cli) CheckHealth() (string, error) {
	return "", nil
}

func (c cli) Topics() ([]string, error) {
	return nil, nil
}

func (c cli) Publish(topic string, message []byte) error {
	return nil
}

func (c cli) Subscribe(topic string) chan []byte {
	return nil
}
