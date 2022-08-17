//go:build !linux && !windows
// +build !linux,!windows

package service // import "github.com/fdurand/moby/volume/service"

import (
	"github.com/fdurand/moby/pkg/idtools"
	"github.com/fdurand/moby/volume/drivers"
)

func setupDefaultDriver(_ *drivers.Store, _ string, _ idtools.Identity) error { return nil }
