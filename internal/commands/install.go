// internal/commands/install.go
package commands

import (
	"fmt"
	"os"
	"os/exec"
)

func Install(args []string) {
	if len(args) < 1 {
		fmt.Println("You must specify a package to install")
		return
	}

	packageName := args[0]
	fmt.Printf("Installing package: %s\n", packageName)

	cmd := exec.Command("go", "get", packageName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error installing package: %v\n", err)
	}
}
