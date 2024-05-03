package utils

import "github.com/schollz/progressbar/v3"

// NewProgressBar generates a new schollz progressbar in order to display
// the publish and subscribe statsu
func NewProgressBar(max int, name string) *progressbar.ProgressBar {
	return progressbar.Default(int64(max), name)
}
