package packageinfo

import (
	"errors"
	"fmt"
	"os/exec"
)

// PackageInfo holds information about a Go package.
type PackageInfo struct {
	Name    string
	Version string
}

// GetPackageInfo retrieves information about a Go package.
func GetPackageInfo(packageName string) (*PackageInfo, error) {
	cmd := exec.Command("go", "list", "-m", "-json", packageName)
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	// Parse the JSON output and populate PackageInfo struct
	// For simplicity, we'll assume the output is just the package name and version
	// In a real implementation, you would use a JSON library to parse the output
	info := &PackageInfo{Name: packageName, Version: "v0.0.0"} // Placeholder version
	fmt.Printf("Package Info: %s\n", output)
	return info, nil
}
