package utils

import (
	"encoding/json"
	"os"
)

// WriteJSON is used to export any type to a json file
func WriteJSON[T any](path string, instance *T) error {
	bytes, err := json.Marshal(instance)
	if err != nil {
		return err
	}

	if err := os.WriteFile(path, bytes, os.ModeAppend); err != nil {
		return err
	}

	return nil
}

// ReadJSON is used to read a json file to a any type
func ReadJSON[T any](path string) (*T, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var instance T

	if err := json.Unmarshal(bytes, &instance); err != nil {
		return nil, err
	}

	return &instance, nil
}
