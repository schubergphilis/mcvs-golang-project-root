package projectroot

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

// FindProjectRoot finds the root directory of the project.
func FindProjectRoot() (string, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("unable to return path for the corresponding current working dir: %w", err)
	}

	for {
		if fileExists(filepath.Join(currentDir, "go.mod")) {
			return currentDir, nil
		}

		parentDir := filepath.Dir(currentDir)
		if currentDir == parentDir {
			break
		}

		currentDir = parentDir
	}

	return "", errors.New("project root not found")
}

func fileExists(path string) bool {
	_, err := os.Stat(path)

	return !os.IsNotExist(err)
}
