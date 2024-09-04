package user

import (
	"os"
	"testing"
)

func TestSetupUserNamespace(t *testing.T) {
	if os.Geteuid() != 0 {
		t.Skip("Test requires root privileges")
	}

	cmd := "/bin/sh"
	args := []string{"-c", "echo 'User namespace works!'"}

	err := SetupUserNamespace(cmd, args)
	if err != nil {
		t.Fatalf("failed to setup user namespace: %v", err)
	}

	t.Log("User namespace setup successfully")
}
