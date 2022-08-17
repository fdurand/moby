//go:build windows
// +build windows

package runconfig // import "github.com/fdurand/moby/runconfig"

import (
	"testing"

	"github.com/fdurand/moby/api/types/container"
)

func TestValidatePrivileged(t *testing.T) {
	expected := "Windows does not support privileged mode"
	err := validatePrivileged(&container.HostConfig{Privileged: true})
	if err == nil || err.Error() != expected {
		t.Fatalf("Expected %s", expected)
	}
}
