package mqtt

import "errors"

var (
	ErrConnection = errors.New("failed to connect to the cluster")
)
