package uts

import (
	"fmt"
	"os"
	"syscall"
)

func SetupUTSNamespace(cmd string, args []string, hostname string) error {
	procAttr := syscall.ProcAttr{
		Env:   os.Environ(),
		Files: []uintptr{0, 1, 2},
		Sys: &syscall.SysProcAttr{
			Cloneflags: syscall.CLONE_NEWUTS,
		},
	}

	pid, err := syscall.ForkExec(cmd, args, &procAttr)
	if err != nil {
		return fmt.Errorf("failed to fork and exec: %v", err)
	}

	_, err = syscall.Wait4(pid, nil, 0, nil)
	if err != nil {
		return fmt.Errorf("failed to wait for child process: %v", err)
	}

	if err = syscall.Sethostname([]byte(hostname)); err != nil {
		return fmt.Errorf("failed to set hostname: %v", err)
	}

	return nil
}
