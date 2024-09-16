package composeclient

import (
	"context"
	"strings"

	"github.com/compose-spec/compose-go/v2/cli"
	"github.com/docker/cli/cli/command"
	"github.com/docker/cli/cli/flags"
	commands "github.com/docker/compose/v2/cmd/compose"
	"github.com/docker/compose/v2/pkg/api"
	"github.com/docker/compose/v2/pkg/compose"
)

func Up(content []byte, projectName string) error {
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

	for name, s := range project.Services {
		s.CustomLabels = map[string]string{
			api.ProjectLabel:     project.Name,
			api.ServiceLabel:     name,
			api.VersionLabel:     api.ComposeVersion,
			api.WorkingDirLabel:  project.WorkingDir,
			api.ConfigFilesLabel: strings.Join(project.ComposeFiles, ","),
			api.OneoffLabel:      "False", // default, will be overridden by `run` command
		}
		project.Services[name] = s
	}

	buildOption := api.BuildOptions{
		Pull:     true,
		Push:     false,
		Progress: "auto",
		NoCache:  false,
		Quiet:    true,
		Services: project.ServiceNames(),
		Deps:     false,
	}

	createOption := api.CreateOptions{
		Build:         &buildOption,
		Services:      project.ServiceNames(),
		RemoveOrphans: false,
		IgnoreOrphans: false,
		Recreate:      api.RecreateNever,
		Inherit:       false,
		QuietPull:     true,
	}

	startOption := api.StartOptions{
		Project:  project,
		OnExit:   api.CascadeIgnore,
		Services: project.ServiceNames(),
	}

	upOption := api.UpOptions{
		Create: createOption,
		Start:  startOption,
	}
	err = backend.Up(ctx, project, upOption)

	return err
}
