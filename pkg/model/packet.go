package model

import "time"

type Packet struct {
	Timestamp time.Time
	Payload   []byte
}
