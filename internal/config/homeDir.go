package config

import (
	"fmt"
	"os"
	"path/filepath"
)

func FilePath(path string) (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to get user home directory: %v", err)
	}
	ProjectFir := filepath.Join(homeDir, "workspace/github.com/felixsolom/gator/")
	fPath := filepath.Join(ProjectFir, path)
	return fPath, nil
}
