package composeclient

import (
	"context"

	"github.com/docker/cli/cli/command"
	"github.com/docker/cli/cli/flags"
	"github.com/docker/compose/v2/pkg/api"
	moby "github.com/docker/docker/api/types"
	containerType "github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
)

func List() ([]moby.Container, error) {
	dockerCli, err := command.NewDockerCli()
	if err != nil {
		return make([]moby.Container, 0), err
	}

	clientOptions := flags.NewClientOptions()
	clientOptions.Context = dockerCli.CurrentContext()
	err = dockerCli.Initialize(clientOptions)
	if err != nil {
		return make([]moby.Container, 0), err
	}

	return dockerCli.Client().ContainerList(context.Background(), containerType.ListOptions{
		Filters: filters.NewArgs(hasProjectLabelFilter(), hasComposeCLientLabelFilter()),
		All:     true,
	})
}

func hasProjectLabelFilter() filters.KeyValuePair {
	return filters.Arg("label", api.ProjectLabel)
}

func hasComposeCLientLabelFilter() filters.KeyValuePair {
	return filters.Arg("label", COMPOSE_CLIENT_LABEL)
}
