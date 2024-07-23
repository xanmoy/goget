package utils

import (
	"testing"
)

func TestRunCommand(t *testing.T) {
	output, err := RunCommand("echo", "hello")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := "hello\n"
	if output != expected {
		t.Errorf("Expected output to be %q, got %q", expected, output)
	}
}
