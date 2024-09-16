package composeclient

import (
	"os"
	"path/filepath"
)

func createComposeFile(content []byte, projectName string) (string, error) {
	dirName := filepath.Join("compose", projectName)

	os.RemoveAll(dirName)
	os.MkdirAll(dirName, 0644)

	composeFile := filepath.Join(dirName, "docker-compose.yml")

	err := os.WriteFile(composeFile, content, 0644)

	return composeFile, err
}
