package config

import (
	"fmt"

	"github.com/amirhnajafiz/cemq/internal/utils"
)

type config struct{}

// Select method is used to change the current context
func Select(name string) string {
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
