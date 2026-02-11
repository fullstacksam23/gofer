package tools

import (
	"bytes"
	"os"
	"os/exec"
)

func BashTool(command string) (string, error) {
	var cmd *exec.Cmd

	if os.Getenv("OS") == "Windows_NT" {
		cmd = exec.Command("cmd", "/C", command)
	} else {
		cmd = exec.Command("sh", "-c", command)
	}

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	output := stdout.String() + stderr.String()

	if err != nil {
		return output, err
	}

	return output, nil
}
