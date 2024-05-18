package config

import (
	"fmt"

	"github.com/amirhnajafiz/cemq/internal/utils"
)

const (
	ROOT    = "~/.config"
	BASE    = "cemq"
	CONFIGS = "configs"
	CONTEXT = "config.txt"
)

// setupBaseConfigDirectories makes the following entities
// .config/
// .config/cemq/
// .conifg/cemq/configs/
// .config/cemq/config.txt
func setupBaseConfigDirectories() {
	if ok, err := utils.Exists(ROOT); err == nil && !ok {
		utils.Mkdir(ROOT)
	}

	baseDirectory := fmt.Sprintf("%s/%s", ROOT, BASE)
	if ok, err := utils.Exists(baseDirectory); err == nil && !ok {
		utils.Mkdir(baseDirectory)
	}

	configsDirectory := fmt.Sprintf("%s/%s", baseDirectory, CONFIGS)
	if ok, err := utils.Exists(configsDirectory); err == nil && !ok {
		utils.Mkdir(configsDirectory)
	}
}
