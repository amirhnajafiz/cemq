package utils

import "os"

// ReadFile returns the string value of a file
func ReadFile(path string) (string, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

// WriteFile creates a new file with given content
func WriteFile(path, content string) error {
	if err := os.WriteFile(path, []byte(content), os.FileMode(os.O_RDWR)); err != nil {
		return err
	}

	return nil
}

// Mkdir is used to created a directory
func Mkdir(path string) {
	os.Mkdir(path, os.ModeDir)
}

// Exists check the existance of a file or folder
func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}

	return false, err
}
