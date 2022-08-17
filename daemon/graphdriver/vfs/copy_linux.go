package vfs // import "github.com/fdurand/moby/daemon/graphdriver/vfs"

import "github.com/fdurand/moby/daemon/graphdriver/copy"

func dirCopy(srcDir, dstDir string) error {
	return copy.DirCopy(srcDir, dstDir, copy.Content, false)
}
