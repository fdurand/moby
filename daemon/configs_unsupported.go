//go:build !linux && !windows
// +build !linux,!windows

package daemon // import "github.com/fdurand/moby/daemon"

func configsSupported() bool {
	return false
}
