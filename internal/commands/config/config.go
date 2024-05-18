package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/amirhnajafiz/cemq/internal/utils"
	"github.com/amirhnajafiz/cemq/pkg/adr"
)

type config struct{}

// List returns the list of config files
func (c config) List() []string {
	fileList := []string{}

	if err := filepath.Walk(adr.GetConfigs(), func(path string, f os.FileInfo, err error) error {
		fileList = append(fileList, path)

		return nil
	}); err != nil {
		return fileList
	}

	return fileList
}

// Info reads the current config file
func (c config) Info() string {
	obj, err := adr.LoadCurrentConfig()
	if err != nil {
		return fmt.Errorf("failed to read current config file: %w", err).Error()
	}

	return fmt.Sprintf("%v", obj)
}

// Select method is used to change the current context
func (c config) Select(name string) string {
	path := adr.GetContext()
	out := fmt.Sprintf("%s/%s.json", adr.GetConfigs(), name)

	if ok, err := utils.Exists(out); err == nil && !ok {
		return fmt.Errorf("context %s not found", err).Error()
	}

	if err := utils.WriteFile(path, out); err != nil {
		return fmt.Errorf("failed to select %s context: %w", name, err).Error()
	}

	return fmt.Sprintf("selected: %s", name)
}
