package projectroot

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

// FindProjectRoot finds the root directory of the project by walking up from
// the working directory to the nearest directory containing a go.mod file.
// The result is not a trust boundary: do not rely on it for security-sensitive
// file access.
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
	info, err := os.Stat(path)

	return err == nil && info.Mode().IsRegular()
}
