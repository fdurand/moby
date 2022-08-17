package daemon // import "github.com/fdurand/moby/daemon"

import (
	"github.com/fdurand/moby/api/types/container"
	libcontainerdtypes "github.com/fdurand/moby/libcontainerd/types"
)

func toContainerdResources(resources container.Resources) *libcontainerdtypes.Resources {
	// We don't support update, so do nothing
	return nil
}
