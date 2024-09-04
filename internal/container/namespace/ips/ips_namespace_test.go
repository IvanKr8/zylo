package ips

import (
	"os"
	"testing"
)

func TestSetupIPCNamespace(t *testing.T) {
	testDir := "/tmp/test_ipc_namespace"
	if err := os.Mkdir(testDir, 0755); err != nil && !os.IsExist(err) {
		t.Fatalf("failed to create test directory: %v", err)
	}
	defer os.RemoveAll(testDir)

	cmd := "/bin/sh"
	args := []string{"-c", "sleep 10"}

	if err := setupIPCNamespace(cmd, args); err != nil {
		t.Fatalf("setupIPCNamespace failed: %v", err)
	}
}
