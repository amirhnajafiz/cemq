package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/amirhnajafiz/cemq/internal/utils"
	"github.com/amirhnajafiz/cemq/pkg/adr"
)

type config struct{}

// List returns the list of config files
func (c config) List() []string {
	fileList := []string{}

	if err := filepath.Walk(adr.GetConfigs(), func(path string, f os.FileInfo, err error) error {
		name := f.Name()
		if name == "configs" {
			return nil
		}

		fileList = append(fileList, strings.Replace(name, ".json", "", 1))

		return nil
	}); err != nil {
		return fileList
	}

	return fileList
}

// Info reads the current config file
func (c config) Info() string {
	path, err := utils.ReadFile(adr.GetContext())
	if err != nil {
		return fmt.Errorf("failed to read config.txt file: %w", err).Error()
	}

	data, err := utils.ReadFile(path)
	if err != nil {
		return fmt.Errorf("failed to read current config file: %w", err).Error()
	}

	return utils.JsonPrettyPrint(data)
}

// Select method is used to change the current context
func (c config) Select(name string) string {
	path := adr.GetContext()
	out := fmt.Sprintf("%s/%s.json", adr.GetConfigs(), name)

	if ok, err := utils.Exists(out); err == nil && !ok {
		return fmt.Errorf("context %s not found", name).Error()
	}

	if err := utils.WriteFile(path, out); err != nil {
		return fmt.Errorf("failed to select %s context: %w", name, err).Error()
	}

	return fmt.Sprintf("selected: %s", name)
}
