package user

import (
	"fmt"
	"os"
	"syscall"
)

func SetupUserNamespace(cmd string, args []string) error {
	procAttr := syscall.ProcAttr{
		Env:   os.Environ(),
		Files: []uintptr{0, 1, 2},
		Sys: &syscall.SysProcAttr{
			Cloneflags: syscall.CLONE_NEWUSER,
			UidMappings: []syscall.SysProcIDMap{
				{
					ContainerID: 0,
					HostID:      os.Getuid(),
					Size:        1,
				},
			},
			GidMappings: []syscall.SysProcIDMap{
				{
					ContainerID: 0,
					HostID:      os.Getgid(),
					Size:        1,
				},
			},
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

	return nil
}
