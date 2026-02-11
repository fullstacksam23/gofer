package tools

import (
	"os"
	"path/filepath"
)

func WriteFileTool(filePath, content string) error {
	err := os.MkdirAll(filepath.Dir(filePath), 0755)
	if err != nil {
		return err
	}
	return os.WriteFile(filePath, []byte(content), 0644)
}
