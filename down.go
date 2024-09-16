package composeclient

import (
	"context"

	"github.com/compose-spec/compose-go/v2/cli"
	"github.com/docker/cli/cli/command"
	"github.com/docker/cli/cli/flags"
	commands "github.com/docker/compose/v2/cmd/compose"
	"github.com/docker/compose/v2/pkg/api"
	"github.com/docker/compose/v2/pkg/compose"
)

func Down(content []byte, projectName string) error {
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

	project, err := cli.ProjectFromOptions(ctx, options)
	if err != nil {
		return err
	}

	dockerCli, err := command.NewDockerCli()
	if err != nil {
		return err
	}

	clientOptions := flags.NewClientOptions()
	clientOptions.Context = dockerCli.CurrentContext()
	err = dockerCli.Initialize(clientOptions)
	if err != nil {
		return err
	}

	backend := compose.NewComposeService(dockerCli).(commands.Backend)

	return backend.Down(ctx, projectName, api.DownOptions{
		RemoveOrphans: false,
		Project:       project,
		Services:      project.ServiceNames(),
	})
}
