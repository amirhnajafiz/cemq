package utils

import "sync"

// SafeCounter is safe to use concurrently
type SafeCounter struct {
	channel chan bool
	expired bool
	mu      sync.Mutex
	v       int
	l       int
}

func NewSafeCounter(limit int) *SafeCounter {
	return &SafeCounter{
		channel: make(chan bool),
		v:       0,
		l:       limit,
		expired: false,
	}
}

// Inc increments the counter
func (c *SafeCounter) Inc() {
	if c.expired {
		return
	}

	c.mu.Lock()

	c.v++

	if c.v == c.l {
		c.channel <- true
	}

	c.mu.Unlock()
}

// Wait will stop the process until it gets a value from check method
func (c *SafeCounter) Wait() {
	c.expired = <-c.channel
}
