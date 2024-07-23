// internal/commands/update.go
package commands

import (
	"fmt"
	"os"
	"os/exec"
)

func Update(args []string) {
	if len(args) < 1 {
		fmt.Println("You must specify a package to update")
		return
	}

	packageName := args[0]
	fmt.Printf("Updating package: %s\n", packageName)

	cmd := exec.Command("go", "get", "-u", packageName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error updating package: %v\n", err)
	}
}
