package utils

import "os"

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
