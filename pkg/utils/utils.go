package utils

import (
	"fmt"
	"os"
	"os/exec"
)

// RunCommand runs a shell command and returns the output or an error.
func RunCommand(name string, args ...string) (string, error) {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(output), nil
}

// PrintError prints an error message to the console.
func PrintError(err error) {
	fmt.Printf("Error: %v\n", err)
}
