package commands

import (
	"fmt"
	"github.com/IvanKr8/zylo/internal/zyloFile/types"
	"os"
	"os/exec"
	"path/filepath"
)

func ExecuteCommands(config *types.ZyloConfig) error {
	if config.Workdir != "" {
		fmt.Printf("Setting working directory to %s...\n", config.Workdir)
		if err := os.MkdirAll(config.Workdir, os.ModePerm); err != nil {
			return fmt.Errorf("failed to create workdir: %v", err)
		}
	}

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

	for _, command := range config.Commands {
		fmt.Printf("Executing command: %s\n", command)
		cmd := exec.Command("sh", "-c", command)
		cmd.Dir = config.Workdir
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("failed to execute command %s: %v", command, err)
		}
	}

	if config.Entrypoint != "" {
		fmt.Printf("Starting entrypoint: %s\n", config.Entrypoint)
		cmd := exec.Command("sh", "-c", config.Entrypoint)
		cmd.Dir = config.Workdir
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("failed to start entrypoint %s: %v", config.Entrypoint, err)
		}
	}

	return nil
}
