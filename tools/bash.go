package tools

import (
	"bytes"
	"os/exec"
)

func BashTool(command string) (string, error) {
	cmd := exec.Command("sh", "-c", command)

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
