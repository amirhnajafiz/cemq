package config

import (
	"fmt"

	"github.com/amirhnajafiz/cemq/internal/utils"
)

const (
	ORGPATH = "~/.config"
	SUBDIR  = "cemq"
	CONFIGS = "configs"
	CONTEXT = "config.txt"
)

func init() {
	base := fmt.Sprintf("%s/%s", ORGPATH, SUBDIR)
	configs := fmt.Sprintf("%s/%s", base, CONFIGS)

	if ok, err := utils.Exists(ORGPATH); err == nil && !ok {
		utils.Mkdir(ORGPATH)
	}

	if ok, err := utils.Exists(base); err == nil && !ok {
		utils.Mkdir(base)
	}

	if ok, err := utils.Exists(configs); err == nil && !ok {
		utils.Mkdir(configs)
	}
}
