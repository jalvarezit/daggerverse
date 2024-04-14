package main

import (
	"context"
)

type ContainerStructureTest struct{}

func (m *ContainerStructureTest) Test(ctx context.Context, image string, config *File) *Container {
	podmanService := dag.Podman().TCPService(PodmanTCPServiceOpts{Port: 1337, Image: []string{image}})

	configFile, err := config.Name(ctx)

	if err != nil {
		panic(err)
	}

	return dag.Container().
		From("gcr.io/gcp-runtimes/container-structure-test").
		WithServiceBinding("podman", podmanService).
		WithEnvVariable("DOCKER_HOST", "tcp://podman:1337").
		WithFile(configFile, config).
		WithExec([]string{"test", "-i", image, "--config", configFile})
}
