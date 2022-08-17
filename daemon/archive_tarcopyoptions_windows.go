package daemon // import "github.com/fdurand/moby/daemon"

import (
	"github.com/fdurand/moby/container"
	"github.com/fdurand/moby/pkg/archive"
)

func (daemon *Daemon) tarCopyOptions(container *container.Container, noOverwriteDirNonDir bool) (*archive.TarOptions, error) {
	return daemon.defaultTarCopyOptions(noOverwriteDirNonDir), nil
}
