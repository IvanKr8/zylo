package mount

import (
	"fmt"
	"os"
	"syscall"
)

func setupMountNamespace(cmd string, args []string, chRootDir string) error {
	procAttr := syscall.ProcAttr{
		Env:   os.Environ(),
		Files: []uintptr{0, 1, 2},
		Sys: &syscall.SysProcAttr{
			Cloneflags: syscall.CLONE_NEWNS,
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

	if err = syscall.Chroot(chRootDir); err != nil {
		return fmt.Errorf("failed to chroot: %v", err)
	}

	return nil
}
