package config

import (
	"github.com/amirhnajafiz/cemq/internal/utils"
	"github.com/amirhnajafiz/cemq/pkg/adr"
)

// setupBaseConfigDirectories makes the following entities
// .config/
// .config/cemq/
// .conifg/cemq/configs/
func setupBaseConfigDirectories() {
	if ok, err := utils.Exists(adr.GetRoot()); err == nil && !ok {
		utils.Mkdir(adr.GetRoot())
	}

	if ok, err := utils.Exists(adr.GetBase()); err == nil && !ok {
		utils.Mkdir(adr.GetBase())
	}

	if ok, err := utils.Exists(adr.GetConfigs()); err == nil && !ok {
		utils.Mkdir(adr.GetConfigs())
	}
}
