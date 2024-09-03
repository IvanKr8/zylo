package mount

import (
	"os"
	"testing"
	"time"
)

func TestSetupMountNamespace(t *testing.T) {
	// Define the test directory
	testDir := "/tmp/test_chroot"

	// Remove the directory if it already exists
	if _, err := os.Stat(testDir); !os.IsNotExist(err) {
		if err = os.RemoveAll(testDir); err != nil {
			t.Fatalf("failed to clean up existing test directory: %v", err)
		}
	}

	// Create a new directory for the test
	if err := os.Mkdir(testDir, 0755); err != nil {
		t.Fatalf("failed to create test directory: %v", err)
	}
	defer os.RemoveAll(testDir) // Clean up after the test

	// Create a dummy file in the test directory
	testFile := testDir + "/testfile"
	file, err := os.Create(testFile)
	if err != nil {
		t.Fatalf("failed to create test file: %v", err)
	}
	file.Close()

	// Simple command to run in the new namespace
	cmd := "/bin/sh"
	args := []string{"-c", "sleep 10"} // Short sleep to speed up the test

	// Execute setupMountNamespace
	if err = setupMountNamespace(cmd, args, testDir); err != nil {
		t.Fatalf("setupMountNamespace failed: %v", err)
	}

	// Allow some time for the process to complete
	time.Sleep(1 * time.Second)

	// Check if the test file is still accessible, indicating successful chroot
	if _, err = os.Stat(testFile); !os.IsNotExist(err) {
		t.Fatalf("test file should not be visible outside the chroot directory")
	}
}
