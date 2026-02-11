package tools

import "os"

func ReadFileTool(filePath string) (string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	content := string(data)
	return content, nil
}
