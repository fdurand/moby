//go:build !linux
// +build !linux

package gcplogs // import "github.com/fdurand/moby/daemon/logger/gcplogs"

func ensureHomeIfIAmStatic() error {
	return nil
}
