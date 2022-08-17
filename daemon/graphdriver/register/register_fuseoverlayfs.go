//go:build !exclude_graphdriver_fuseoverlayfs && linux
// +build !exclude_graphdriver_fuseoverlayfs,linux

package register // import "github.com/fdurand/moby/daemon/graphdriver/register"

import (
	// register the fuse-overlayfs graphdriver
	_ "github.com/fdurand/moby/daemon/graphdriver/fuse-overlayfs"
)
