package engine

import (
	"github.com/IvanKr8/zylo/internal/zyloFile/types"
)

func CreateContainer(config *types.ZyloConfig) error {
	if err := createWorkdir(config.Workdir); err != nil {
		return err
	}

	if err := copyFiles(config); err != nil {
		return err
	}

	if err := executeCommands(config.Commands, config.Workdir); err != nil {
		return err
	}

	if err := startEntrypoint(config.Entrypoint, config.Workdir); err != nil {
		return err
	}

	return nil
}
