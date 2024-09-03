package pid

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func SetupTestPIDNamespace(t *testing.T) {
	cmd := "/bin/sh"
	args := []string{"-c", "echo $$ > /tmp/pid.txt && sleep 10"}

	err := SetupPIDNamespace(cmd, args)
	assert.NoError(t, err, "Error should be nil")

	_, err = os.Stat("/tmp/pid.txt")
	assert.NoError(t, err, "PID file should exist")

	err = os.Remove("/tmp/pid.txt")
	assert.NoError(t, err, "Error removing PID file should be nil")
}
