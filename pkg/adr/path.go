package adr

import "fmt"

const (
	root    = "~/.config"
	base    = "cemq"
	configs = "configs"
	context = "config.txt"
)

func GetRoot() string {
	return root
}

func GetBase() string {
	return fmt.Sprintf("%s/%s", root, base)
}

func GetContext() string {
	return fmt.Sprintf("%s/%s/%s", root, base, context)
}

func GetConfigs() string {
	return fmt.Sprintf("%s/%s/%s", root, base, configs)
}
