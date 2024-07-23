// internal/commands/remove.go
package commands

import (
    "bufio"
    "fmt"
    "os"
    "os/exec"
    "strings"
)

func Remove(args []string) {
    if len(args) < 1 {
        fmt.Println("You must specify a package to remove")
        return
    }

    packageName := args[0]
    fmt.Printf("Removing package: %s\n", packageName)

    err := removePackageFromGoMod(packageName)
    if err != nil {
        fmt.Printf("Error removing package: %v\n", err)
        return
    }

    err = tidyGoMod()
    if err != nil {
        fmt.Printf("Error tidying go.mod: %v\n", err)
        return
    }

    fmt.Printf("Package %s removed successfully\n", packageName)
}

func removePackageFromGoMod(packageName string) error {
    goModFile, err := os.Open("go.mod")
    if err != nil {
        return err
    }
    defer goModFile.Close()

    var lines []string
    scanner := bufio.NewScanner(goModFile)
    for scanner.Scan() {
        line := scanner.Text()
        if !strings.Contains(line, packageName) {
            lines = append(lines, line)
        }
    }

    if err := scanner.Err(); err != nil {
        return err
    }

    err = os.WriteFile("go.mod", []byte(strings.Join(lines, "\n")), 0644)
    if err != nil {
        return err
    }

    return nil
}

func tidyGoMod() error {
    cmd := exec.Command("go", "mod", "tidy")
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    return cmd.Run()
}
