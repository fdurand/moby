//go:build !linux && !windows
// +build !linux,!windows

package system // import "github.com/fdurand/moby/pkg/system"

// ReadMemInfo is not supported on platforms other than linux and windows.
func ReadMemInfo() (*MemInfo, error) {
	return nil, ErrNotSupportedPlatform
}
