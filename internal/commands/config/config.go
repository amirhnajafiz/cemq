package config

import (
	"fmt"
	"os"
)

type config struct{}

// Select method is used to change the current context
func Select(name string) string {
	path := fmt.Sprintf("%s/%s/%s", ROOT, BASE, CONTEXT)
	out := fmt.Sprintf("%s/%s/%s/%s.json", ROOT, BASE, CONFIGS, name)

	if err := os.WriteFile(path, []byte(out), os.FileMode(os.O_RDWR)); err != nil {
		return fmt.Errorf("failed to select context: %w", err).Error()
	}

	return fmt.Sprintf("selected: %s", name)
}
