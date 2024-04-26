package main

import "fmt"

type Podman struct{}

// Returns a service with the podman engine installed and running
func (m *Podman) TcpService(
	// The port to expose the podman service on
	// +optional
	// +default=1337
	port int,

	// Podman image to use
	// +optional
	// +default="quay.io/podman/stable"
	image string,

	// Podman tag to use
	// +optional
	// +default="latest"
	tag string,

	// List of images to pull
	// +optional
	pullImage []string,
) *Service {

	ctr := dag.Container().From(fmt.Sprintf("%s:%s", image, tag))

	// Install images
	for _, img := range pullImage {
		ctr = ctr.WithExec([]string{"podman", "pull", img}, ContainerWithExecOpts{InsecureRootCapabilities: true})
	}

	return ctr.WithExec([]string{"podman", "system", "service", fmt.Sprintf("tcp://0.0.0.0:%d", port), "--time", "0"}, ContainerWithExecOpts{InsecureRootCapabilities: true}).WithExposedPort(port).AsService()

}
