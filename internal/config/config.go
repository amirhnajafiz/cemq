package config

import (
	"github.com/amirhnajafiz/cemq/internal/utils"
	"github.com/amirhnajafiz/cemq/pkg/adr"
	"github.com/amirhnajafiz/cemq/pkg/model"
)

// Load returns the current context as a config struct
func Load() (*model.Config, error) {
	path, err := utils.ReadFile(adr.GetContext())
	if err != nil {
		return nil, err
	}

	obj, err := utils.ReadJSON[model.Config](path)
	if err != nil {
		return nil, err
	}

	return obj, nil
}
