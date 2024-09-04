package cgroup

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

func SetupCgroup(name string, cpuLimit int, pid int) error {
	cgroupPath := filepath.Join("/sys/fs/cgroup/cpu", name)

	if err := os.MkdirAll(cgroupPath, 0755); err != nil {
		return fmt.Errorf("failed to create cgroup directory: %v", err)
	}

	cpuQuotaPath := filepath.Join(cgroupPath, "cpu.cfs_quota_us")
	if err := os.WriteFile(cpuQuotaPath, []byte(strconv.Itoa(cpuLimit)), 0644); err != nil {
		return fmt.Errorf("failed to set CPU quota: %v", err)
	}

	tasksPath := filepath.Join(cgroupPath, "tasks")
	if err := os.WriteFile(tasksPath, []byte(strconv.Itoa(pid)), 0644); err != nil {
		return fmt.Errorf("failed to add process to cgroup: %v", err)
	}

	return nil
}
