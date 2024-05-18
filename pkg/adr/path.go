package adr

import "fmt"

// base variables
const (
	root    = "~/.config"
	base    = "cemq"
	configs = "configs"
	context = "config.txt"
)

// GetRoot returns the root directory
func GetRoot() string {
	return root
}

// GetBase returns CEMQ base directory
func GetBase() string {
	return fmt.Sprintf("%s/%s", root, base)
}

// GetContext returns the path of CEMQ current context
func GetContext() string {
	return fmt.Sprintf("%s/%s/%s", root, base, context)
}

// GetConfigs returns the path of configs directory
func GetConfigs() string {
	return fmt.Sprintf("%s/%s/%s", root, base, configs)
}
