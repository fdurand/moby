package archive // import "github.com/fdurand/moby/pkg/archive"

import (
	"os"

	"github.com/fdurand/moby/pkg/system"
)

func statDifferent(oldStat *system.StatT, newStat *system.StatT) bool {
	// Note there is slight difference between the Linux and Windows
	// implementations here. Due to https://github.com/fdurand/moby/issues/9874,
	// and the fix at https://github.com/fdurand/moby/pull/11422, Linux does not
	// consider a change to the directory time as a change. Windows on NTFS
	// does. See https://github.com/fdurand/moby/pull/37982 for more information.

	if !sameFsTime(oldStat.Mtim(), newStat.Mtim()) ||
		oldStat.Mode() != newStat.Mode() ||
		oldStat.Size() != newStat.Size() && !oldStat.Mode().IsDir() {
		return true
	}
	return false
}

func (info *FileInfo) isDir() bool {
	return info.parent == nil || info.stat.Mode().IsDir()
}

func getIno(fi os.FileInfo) (inode uint64) {
	return
}

func hasHardlinks(fi os.FileInfo) bool {
	return false
}
