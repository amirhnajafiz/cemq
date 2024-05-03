package config

import (
	"fmt"
	"os"
)

type config struct {
	DirPath string
}

func (c config) Select(name string) string {
	path := fmt.Sprintf("%s/%s.json", c.DirPath, name)

	os.Setenv("CEMQ_CONTEXT", path)
}
