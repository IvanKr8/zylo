package pid

import (
	"fmt"
	"os"
	"syscall"
)

func SetupPIDNamespace(cmd string, args []string) error {
	procAttr := syscall.ProcAttr{
		Env:   os.Environ(),
		Files: []uintptr{0, 1, 2},
		Sys: &syscall.SysProcAttr{
			Cloneflags: syscall.CLONE_NEWPID,
		},
	}

	pid, err := syscall.ForkExec(cmd, append([]string{cmd}, args...), &procAttr)
	if err != nil {
		return fmt.Errorf("failed to fork and exec: %v", err)
	}

	_, err = syscall.Wait4(pid, nil, 0, nil)
	if err != nil {
		return fmt.Errorf("failed to wait for child process: %v", err)
	}

	return nil
}
