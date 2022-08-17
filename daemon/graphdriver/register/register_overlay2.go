//go:build !exclude_graphdriver_overlay2 && linux
// +build !exclude_graphdriver_overlay2,linux

package register // import "github.com/fdurand/moby/daemon/graphdriver/register"

import (
	// register the overlay2 graphdriver
	_ "github.com/fdurand/moby/daemon/graphdriver/overlay2"
)
