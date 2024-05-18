package config

import (
	"fmt"
	"os"

	"github.com/amirhnajafiz/cemq/internal/utils"
)

const (
	ROOT    = "~/.config"
	BASE    = "cemq"
	CONFIGS = "configs"
	CONTEXT = "config.txt"
)

// SetupBaseConfigDirectories makes the following entities
// .config/
// .config/cemq/
// .conifg/cemq/configs/
// .config/cemq/config.txt
func SetupBaseConfigDirectories() {
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

// LoadCurrentContext returns the current context path
func LoadCurrentContext() (string, error) {
	path := fmt.Sprintf("%s/%s/%s", ROOT, BASE, CONTEXT)

	bytes, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("failed to load current context: %w", err)
	}

	return string(bytes), nil
}
