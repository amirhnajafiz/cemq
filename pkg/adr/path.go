package adr

import (
	"fmt"
	"os"
)

// base variables
const (
	base    = "cemq"
	configs = "configs"
	context = "config.txt"
)

// GetRoot returns the root directory
func GetRoot() string {
	tmp, _ := os.UserHomeDir()
	return fmt.Sprintf("%s/.config", tmp)
}

// GetBase returns CEMQ base directory
func GetBase() string {
	return fmt.Sprintf("%s/%s", GetRoot(), base)
}

// GetContext returns the path of CEMQ current context
func GetContext() string {
	return fmt.Sprintf("%s/%s/%s", GetRoot(), base, context)
}

// GetConfigs returns the path of configs directory
func GetConfigs() string {
	return fmt.Sprintf("%s/%s/%s", GetRoot(), base, configs)
}
