package config

import (
	"fmt"

	"github.com/amirhnajafiz/cemq/internal/utils"
	"github.com/amirhnajafiz/cemq/pkg/adr"
	"github.com/amirhnajafiz/cemq/pkg/model"
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

		utils.WriteJSON[model.Config](fmt.Sprintf("%s/default.json", adr.GetConfigs()), &model.Config{
			Server:      "tcp://0.0.0.0:1883",
			Description: "Default EMQX cluster",
		})
	}
}
