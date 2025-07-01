package utils

import (
	"os"
	"os/exec"
)

// RunCommand executes shell commands
func RunCommand(cmd string, args ...string) error {
	command := exec.Command(cmd, args...)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	return command.Run()
}
