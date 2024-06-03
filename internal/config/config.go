package config

import (
	"github.com/amirhnajafiz/cemq/internal/utils"
	"github.com/amirhnajafiz/cemq/pkg/adr"
	"github.com/amirhnajafiz/cemq/pkg/model"
)

// Load returns the current context as a config struct
func Load(server string, user string, pass string) (*model.Config, error) {
	if len(server) > 0 {
		return &model.Config{
			Server:   server,
			Username: user,
			Password: pass,
		}, nil
	}

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
