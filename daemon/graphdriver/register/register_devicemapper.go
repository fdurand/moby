//go:build !exclude_graphdriver_devicemapper && !static_build && linux
// +build !exclude_graphdriver_devicemapper,!static_build,linux

package register // import "github.com/fdurand/moby/daemon/graphdriver/register"

import (
	// register the devmapper graphdriver
	_ "github.com/fdurand/moby/daemon/graphdriver/devmapper"
)
