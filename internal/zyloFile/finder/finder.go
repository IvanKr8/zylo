package finder

import (
	"fmt"
	"os"
	"path/filepath"
)

func ZyloFinder(fileName string) (string, error) {
	startDir := "."
	var foundFile string
	err := filepath.Walk(startDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && info.Name() == fileName {
			foundFile = path
			return filepath.SkipDir
		}
		return nil
	})

	if err != nil {
		return "", err
	}

	if foundFile == "" {
		return "", fmt.Errorf("no zylo files found")
	}

	return foundFile, nil
}
