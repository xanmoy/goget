// cmd/goget/main.go
package main

import (
    "fmt"
    "os"
    "github.com/xanmoy/goget/internal/commands"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Expected 'install', 'remove' or 'update' subcommands")
        os.Exit(1)
    }

    switch os.Args[1] {
    case "install":
        commands.Install(os.Args[2:])
    case "remove":
        commands.Remove(os.Args[2:])
    case "update":
        commands.Update(os.Args[2:])
    default:
        fmt.Println("Unknown command")
        os.Exit(1)
    }
}
