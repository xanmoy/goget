package packageinfo

import (
	"testing"
)

func TestGetPackageInfo(t *testing.T) {
	info, err := GetPackageInfo("github.com/sirupsen/logrus")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if info.Name != "github.com/sirupsen/logrus" {
		t.Errorf("Expected package name to be 'github.com/sirupsen/logrus', got %s", info.Name)
	}
}
