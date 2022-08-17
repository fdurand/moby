package libcontainerd // import "github.com/fdurand/moby/libcontainerd"

import (
	"context"

	"github.com/containerd/containerd"
	"github.com/fdurand/moby/libcontainerd/local"
	"github.com/fdurand/moby/libcontainerd/remote"
	libcontainerdtypes "github.com/fdurand/moby/libcontainerd/types"
	"github.com/fdurand/moby/pkg/system"
)

// NewClient creates a new libcontainerd client from a containerd client
func NewClient(ctx context.Context, cli *containerd.Client, stateDir, ns string, b libcontainerdtypes.Backend) (libcontainerdtypes.Client, error) {
	if !system.ContainerdRuntimeSupported() {
		return local.NewClient(ctx, cli, stateDir, ns, b)
	}
	return remote.NewClient(ctx, cli, stateDir, ns, b)
}
