package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/amirhnajafiz/cemq/internal/utils"
	"github.com/amirhnajafiz/cemq/pkg/model"
)

type config struct{}

// List returns the list of config files
func (c config) List() []string {
	root := fmt.Sprintf("%s/%s/%s", ROOT, BASE, CONFIGS)

	fileList := []string{}
	if err := filepath.Walk(root, func(path string, f os.FileInfo, err error) error {
		fileList = append(fileList, path)
		return nil
	}); err != nil {
		return fileList
	}

	return fileList
}

// Info reads the current config file
func (c config) Info() string {
	path := fmt.Sprintf("%s/%s/%s", ROOT, BASE, CONTEXT)

	config, err := utils.ReadFile(path)
	if err != nil {
		return "failed to read current config!"
	}

	obj, err := utils.ReadJSON[*model.Config](config)
	if err != nil {
		return fmt.Errorf("failed to read %s config file: %w", config, err).Error()
	}

	return fmt.Sprintf("%v", obj)
}

// Select method is used to change the current context
func (c config) Select(name string) string {
	path := fmt.Sprintf("%s/%s/%s", ROOT, BASE, CONTEXT)
	out := fmt.Sprintf("%s/%s/%s/%s.json", ROOT, BASE, CONFIGS, name)

	if ok, err := utils.Exists(out); err == nil && !ok {
		return fmt.Errorf("context %s not found", err).Error()
	}

	if err := utils.WriteFile(path, out); err != nil {
		return fmt.Errorf("failed to select %s context: %w", name, err).Error()
	}

	return fmt.Sprintf("selected: %s", name)
}
