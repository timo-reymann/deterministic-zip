package file

import (
	"io/fs"
	"os"
	"path/filepath"
)

// IsDir checks if the given path is a valid directory
func IsDir(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false, err
	}

	return fileInfo.IsDir(), err
}

// ReadDirRecursive reads all files from the given path and returns them always in the same lexical order
func ReadDirRecursive(path string) ([]string, error) {
	paths := make([]string, 0)
	err := filepath.WalkDir(path, func(innerPath string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		paths = append(paths, innerPath)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return paths, nil
}
