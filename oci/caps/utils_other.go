//go:build !linux
// +build !linux

package caps // import "github.com/fdurand/moby/oci/caps"

func initCaps() {
	// no capabilities on Windows
}
