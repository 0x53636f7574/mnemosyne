package utils

import (
	"path/filepath"
	"runtime"
)

func BuildRelativePath(path string) string {
	_, filename, _, _ := runtime.Caller(1)
	return filepath.Join(filepath.Dir(filename), path)
}
