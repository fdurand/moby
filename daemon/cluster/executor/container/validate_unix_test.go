//go:build !windows
// +build !windows

package container // import "github.com/fdurand/moby/daemon/cluster/executor/container"

const (
	testAbsPath        = "/foo"
	testAbsNonExistent = "/some-non-existing-host-path/"
)
