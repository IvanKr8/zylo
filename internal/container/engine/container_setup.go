package engine

import (
	"fmt"
	"github.com/IvanKr8/zylo/internal/zyloFile/types"
	"os"
	"os/exec"
	"path/filepath"
)

func createWorkdir(workdir string) error {
	fmt.Printf("Setting working directory to %s...\n", workdir)
	if err := os.MkdirAll(workdir, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create workdir: %v", err)
	}
	return nil
}

func copyFiles(config *types.ZyloConfig) error {
	for src, dest := range config.Copies {
		fmt.Printf("Copying %s to %s...\n", src, dest)

		if !filepath.IsAbs(dest) {
			dest = filepath.Join(config.Workdir, dest)
		}

		fi, err := os.Stat(src)
		if err != nil {
			return fmt.Errorf("failed to stat source %s: %v", src, err)
		}

		if fi.IsDir() {
			if err = os.MkdirAll(dest, os.ModePerm); err != nil {
				return fmt.Errorf("failed to create destination directory %s: %v", dest, err)
			}
			if err = copyDir(src, dest); err != nil {
				return err
			}
		} else {
			if err = copyFile(src, dest); err != nil {
				return err
			}
		}
	}
	return nil
}

func executeCommands(commands []string, workdir string) error {
	for _, command := range commands {
		fmt.Printf("Executing command: %s\n", command)
		cmd := exec.Command("sh", "-c", command)
		cmd.Dir = workdir
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("failed to execute command %s: %v", command, err)
		}
	}
	return nil
}

func startEntrypoint(entrypoint, workdir string) error {
	if entrypoint != "" {
		fmt.Printf("Starting entrypoint: %s\n", entrypoint)
		cmd := exec.Command("sh", "-c", entrypoint)
		cmd.Dir = workdir
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("failed to start entrypoint %s: %v", entrypoint, err)
		}
	}
	return nil
}
