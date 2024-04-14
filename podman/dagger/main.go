package main

import "fmt"

type Podman struct{}

// Returns a service with the podman engine installed and running
func (m *Podman) TcpService(
	// The port to expose the podman service on
	// +optional
	// +default=1337
	port int,

	// List of images to pull
	// +optional
	image []string,
) *Service {

	ctr := dag.Container().From("quay.io/podman/stable")

	// Install images
	for _, img := range image {
		ctr = ctr.WithExec([]string{"podman", "pull", img}, ContainerWithExecOpts{InsecureRootCapabilities: true})
	}

	return ctr.WithExec([]string{"podman", "system", "service", fmt.Sprintf("tcp://0.0.0.0:%d", port), "--time", "0"}, ContainerWithExecOpts{InsecureRootCapabilities: true}).WithExposedPort(port).AsService()

}
