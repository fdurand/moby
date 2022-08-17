//go:build !linux
// +build !linux

package vfs // import "github.com/fdurand/moby/daemon/graphdriver/vfs"

import (
	"github.com/fdurand/moby/pkg/chrootarchive"
	"github.com/fdurand/moby/pkg/idtools"
)

func dirCopy(srcDir, dstDir string) error {
	return chrootarchive.NewArchiver(idtools.IdentityMapping{}).CopyWithTar(srcDir, dstDir)
}
