package mqtt

import (
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type cli struct {
	conn mqtt.Client
}

func (c cli) CheckConnection() (string, error) {
	token := c.conn.Connect()
	defer func() {
		token.Done()
	}()

	if token.Error() != nil {
		return "", fmt.Errorf("%v: %v", ErrConnection, token.Error())
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
