package commands

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

func copyDir(srcDir string, destDir string) error {
	entries, err := ioutil.ReadDir(srcDir)
	if err != nil {
		return fmt.Errorf("failed to read directory %s: %v", srcDir, err)
	}

	for _, entry := range entries {
		srcPath := filepath.Join(srcDir, entry.Name())
		destPath := filepath.Join(destDir, entry.Name())

		if entry.IsDir() {
			if err = os.MkdirAll(destPath, os.ModePerm); err != nil {
				return fmt.Errorf("failed to create directory %s: %v", destPath, err)
			}

			if err = copyDir(srcPath, destPath); err != nil {
				return err
			}
		} else {
			if err = copyFile(srcPath, destPath); err != nil {
				return err
			}
		}
	}
	return nil
}

func copyFile(srcFile string, destFile string) error {
	input, err := os.Open(srcFile)
	if err != nil {
		return fmt.Errorf("failed to open file %s: %v", srcFile, err)
	}
	defer input.Close()

	output, err := os.Create(destFile)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %v", destFile, err)
	}
	defer output.Close()

	if _, err = io.Copy(output, input); err != nil {
		return fmt.Errorf("failed to copy file %s to %s: %v", srcFile, destFile, err)
	}

	return nil
}
