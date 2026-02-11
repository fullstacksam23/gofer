package tools

import (
	"strings"
	"testing"
)

func TestBashTool(t *testing.T) {
	output, err := BashTool("echo hello")

	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	output = strings.TrimSpace(output)

	if output != "hello" {
		t.Fatalf("expected 'hello', got '%s'", output)
	}
}
func TestBashTool_InvalidCommand(t *testing.T) {
	_, err := BashTool("someinvalidcommand123")

	if err == nil {
		t.Fatal("expected error but got nil")
	}
}

func TestBashToolValidCommand(t *testing.T) {
	_, err := BashTool("echo hello world > a.txt")
	if err != nil {
		t.Fatal("create file test failed")
	}
}
