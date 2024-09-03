package uts

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func TestSetupUTSNamespace(t *testing.T) {
	testDir := "/tmp/test_chroot_" + fmt.Sprintf("%d", time.Now().UnixNano())

	if err := os.Mkdir(testDir, 0755); err != nil {
		t.Fatalf("failed to create test directory: %v", err)
	}
	defer os.RemoveAll(testDir)

	testFile := testDir + "/testfile"
	file, err := os.Create(testFile)
	if err != nil {
		t.Fatalf("failed to create test file: %v", err)
	}
	file.Close()

	cmd := "/bin/sh"
	args := []string{"-c", "sleep 10"}

	hostname := "test-hostname"

	if err = setupUTSNamespace(cmd, args, hostname); err != nil {
		t.Fatalf("setupUTSNamespace failed: %v", err)
	}

	time.Sleep(1 * time.Second)

	if err = checkHostname(hostname); err != nil {
		t.Fatalf("hostname should be %s: %v", hostname, err)
	}
}

func checkHostname(expectedHostname string) error {
	hostname, err := os.Hostname()
	if err != nil {
		return err
	}
	if hostname != expectedHostname {
		return fmt.Errorf("expected hostname %s, got %s", expectedHostname, hostname)
	}
	return nil
}
