package internal

import (
	"golang.org/x/tools/imports"
	"os"
	"path/filepath"
)

func EnsurePathExists(path string) {
	absPath, err := filepath.Abs(path)
	_, err = os.Stat(absPath)
	if err != nil {
		if os.IsNotExist(err) {
			err := os.MkdirAll(path, os.ModePerm)
			if err != nil {
				logger.Fatalf("could not create directory %s: %v", path, err)
			}
		} else {
			logger.Fatalf("could not check if path %s exists: %v", path, err)
		}
	}

}

func FormatSource(filename string, src []byte) ([]byte, error) {
	return imports.Process(filename+".go", src, nil)
}
