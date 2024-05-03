package mqtt

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type cli struct {
	conn mqtt.Client
}

func (c cli) connect() error {
	token := c.conn.Connect()
	defer func() {
		token.Done()
	}()

	if token.Wait() && token.Error() != nil {
		return token.Error()
	}

	return nil
}

func (c cli) Register() error {
	return c.connect()
}

func (c cli) CheckConnection() (string, error) {
	if err := c.connect(); err != nil {
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
