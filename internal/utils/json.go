package utils

import (
	"bytes"
	"encoding/json"
	"os"
)

// WriteJSON is used to export any type to a json file
func WriteJSON[T any](path string, instance *T) error {
	bytes, err := json.Marshal(instance)
	if err != nil {
		return err
	}

	modified := []byte(JsonPrettyPrint(string(bytes)))

	if err := os.WriteFile(path, modified, 0700); err != nil {
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

// JsonPrettyPrint is used to convert JSON raw string to pretty JSON string
func JsonPrettyPrint(in string) string {
	var out bytes.Buffer

	err := json.Indent(&out, []byte(in), "", "  ")
	if err != nil {
		return in
	}

	return out.String()
}
