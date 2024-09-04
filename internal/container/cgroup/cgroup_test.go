package cgroup

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
)

func TestSetupCgroup(t *testing.T) {
	cgroupName := "test_cgroup"
	cpuLimit := 10000
	pid := os.Getpid()

	err := SetupCgroup(cgroupName, cpuLimit, pid)
	if err != nil {
		t.Fatalf("failed to setup cgroup: %v", err)
	}

	cgroupPath := filepath.Join("/sys/fs/cgroup/cpu", cgroupName, "cpu.cfs_quota_us")
	data, err := os.ReadFile(cgroupPath)
	if err != nil {
		t.Fatalf("failed to read cpu.cfs_quota_us: %v", err)
	}

	readLimit, err := strconv.Atoi(string(data))
	if err != nil {
		t.Fatalf("failed to convert cpu limit: %v", err)
	}

	if readLimit != cpuLimit {
		t.Fatalf("expected CPU limit %d, but got %d", cpuLimit, readLimit)
	}

	tasksPath := filepath.Join("/sys/fs/cgroup/cpu", cgroupName, "tasks")
	tasksData, err := os.ReadFile(tasksPath)
	if err != nil {
		t.Fatalf("failed to read tasks file: %v", err)
	}

	pidStr := strconv.Itoa(pid)
	tasks := string(tasksData)
	if !contains(tasks, pidStr) {
		t.Fatalf("process with PID %d not found in cgroup tasks", pid)
	}

	err = os.RemoveAll(filepath.Join("/sys/fs/cgroup/cpu", cgroupName))
	if err != nil {
		t.Fatalf("failed to remove cgroup: %v", err)
	}
}

func contains(tasks string, pidStr string) bool {
	for _, task := range strings.Split(tasks, "\n") {
		if task == pidStr {
			return true
		}
	}
	return false
}
