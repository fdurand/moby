//go:build !linux
// +build !linux

package daemon // import "github.com/fdurand/moby/daemon"

import (
	"context"

	"github.com/containerd/containerd/containers"
	coci "github.com/containerd/containerd/oci"
	"github.com/fdurand/moby/container"
)

const supportsSeccomp = false

// WithSeccomp sets the seccomp profile
func WithSeccomp(daemon *Daemon, c *container.Container) coci.SpecOpts {
	return func(ctx context.Context, _ coci.Client, _ *containers.Container, s *coci.Spec) error {
		return nil
	}
}
