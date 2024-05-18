package mqtt

import (
	"fmt"
	"time"

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
	done := make(chan bool)
	topic := "testing-cluster-health"
	msg := "empty-msg"

	if token := c.conn.Subscribe(topic, 0, func(_ mqtt.Client, m mqtt.Message) {
		m.Ack()

		done <- true
	}); token.Wait() && token.Error() != nil {
		return "", token.Error()
	}

	start := time.Now()

	if token := c.conn.Publish(topic, 0, true, msg); token.Wait() && token.Error() != nil {
		return "", token.Error()
	}

	<-done

	return fmt.Sprintf("ping-pong duration %s", time.Since(start).String()), nil
}

func (c cli) Publish(topic string, message []byte) error {
	return nil
}

func (c cli) Subscribe(topic string) chan []byte {
	return nil
}
