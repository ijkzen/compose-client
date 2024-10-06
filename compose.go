package composeclient

import (
	"context"
	"os"
	"path/filepath"

	"github.com/compose-spec/compose-go/v2/cli"
)

func createComposeFile(content []byte, projectName string) (string, error) {
	dirName := filepath.Join("compose", projectName)

	os.RemoveAll(dirName)
	os.MkdirAll(dirName, 0644)

	composeFile := filepath.Join(dirName, "docker-compose.yml")

	err := os.WriteFile(composeFile, content, 0644)

	return composeFile, err
}

func CheckComposeFileValid(content []byte, projectName string) error {
	composeFilePath, err := createComposeFile(content, projectName)
	if err != nil {
		return err
	}

	ctx := context.Background()

	options, err := cli.NewProjectOptions(
		[]string{composeFilePath},
		cli.WithOsEnv,
		cli.WithDotEnv,
		cli.WithName(projectName),
	)
	if err != nil {
		return err
	}

	_, err = cli.ProjectFromOptions(ctx, options)
	return err
}
