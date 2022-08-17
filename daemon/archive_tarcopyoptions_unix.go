//go:build !windows
// +build !windows

package daemon // import "github.com/fdurand/moby/daemon"

import (
	"github.com/fdurand/moby/container"
	"github.com/fdurand/moby/pkg/archive"
	"github.com/fdurand/moby/pkg/idtools"
)

func (daemon *Daemon) tarCopyOptions(container *container.Container, noOverwriteDirNonDir bool) (*archive.TarOptions, error) {
	if container.Config.User == "" {
		return daemon.defaultTarCopyOptions(noOverwriteDirNonDir), nil
	}

	user, err := idtools.LookupUser(container.Config.User)
	if err != nil {
		return nil, err
	}

	identity := idtools.Identity{UID: user.Uid, GID: user.Gid}

	return &archive.TarOptions{
		NoOverwriteDirNonDir: noOverwriteDirNonDir,
		ChownOpts:            &identity,
	}, nil
}
