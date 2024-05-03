package utils

import "github.com/schollz/progressbar/v3"

func NewProgressBar(max int, name string) *progressbar.ProgressBar {
	return progressbar.Default(int64(max), name)
}
