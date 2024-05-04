package config

import (
	"fmt"
)

type config struct {
	DirPath string
}

func (c config) Select(name string) string {
	path := fmt.Sprintf("%s/config/%s.json", c.DirPath, name)
}
