//go:build !windows
// +build !windows

package dockerfile // import "github.com/fdurand/moby/builder/dockerfile"

func defaultShellForOS(os string) []string {
	return []string{"/bin/sh", "-c"}
}
