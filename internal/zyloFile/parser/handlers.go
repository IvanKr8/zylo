package parser

import (
	"fmt"
	"github.com/IvanKr8/zylo/internal/zyloFile/types"
	"strings"
)

func handleUseImage(args []string, config *types.ZyloConfig) error {
	if len(args) != 1 {
		return fmt.Errorf("USE_IMAGE expects exactly one argument")
	}
	config.Image = args[0]
	fmt.Printf("Using image: %s\n", config.Image)
	return nil
}

func handleSetWorkdir(args []string, config *types.ZyloConfig) error {
	if len(args) != 1 {
		return fmt.Errorf("SET_WORKDIR expects exactly one argument")
	}
	config.Workdir = args[0]
	fmt.Printf("Setting workdir: %s\n", config.Workdir)
	return nil
}

func handleCopyTo(args []string, config *types.ZyloConfig) error {
	if len(args) != 3 || args[1] != "FROM" {
		return fmt.Errorf("COPY_TO expects format 'COPY_TO <dest> FROM <src>'")
	}
	config.Copies[args[2]] = args[0]
	fmt.Printf("Copying %s to %s\n", args[2], args[0])

	return nil
}

func handleExecute(args []string, config *types.ZyloConfig) error {
	command := strings.Join(args, " ")
	config.Commands = append(config.Commands, command)
	fmt.Printf("Adding command to execute: %s\n", command)

	return nil
}

func handleSetEnv(args []string, config *types.ZyloConfig) error {
	if len(args) != 1 || !strings.Contains(args[0], "=") {
		return fmt.Errorf("SET_ENV expects format 'SET_ENV VAR=value'")
	}
	parts := strings.SplitN(args[0], "=", 2)
	config.EnvVars[parts[0]] = parts[1]
	fmt.Printf("Setting environment variable: %s=%s\n", parts[0], parts[1])
	return nil
}

func handleOpenPort(args []string, config *types.ZyloConfig) error {
	if len(args) != 1 {
		return fmt.Errorf("OPEN_PORT expects exactly one argument")
	}
	config.Ports = append(config.Ports, args[0])
	fmt.Printf("Opening port: %s\n", args[0])
	return nil
}

func handleStartWith(args []string, config *types.ZyloConfig) error {
	if len(args) != 1 {
		return fmt.Errorf("START_WITH expects exactly one argument")
	}
	config.Entrypoint = args[0]
	fmt.Printf("Setting entrypoint: %s\n", config.Entrypoint)
	return nil
}
