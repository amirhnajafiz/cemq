package adr

import (
	"github.com/amirhnajafiz/cemq/internal/utils"
	"github.com/amirhnajafiz/cemq/pkg/model"
)

// LoadCurrentConfig is used to read the current config module into a Config struct
// to be used accross CEMQ sub-commands
func LoadCurrentConfig() (*model.Config, error) {
	config, err := utils.ReadFile(GetContext())
	if err != nil {
		return nil, err
	}

	obj, err := utils.ReadJSON[model.Config](config)
	if err != nil {
		return nil, err
	}

	return obj, nil
}
