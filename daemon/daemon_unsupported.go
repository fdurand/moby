//go:build !linux && !freebsd && !windows
// +build !linux,!freebsd,!windows

package daemon // import "github.com/fdurand/moby/daemon"

import (
	"github.com/fdurand/moby/daemon/config"
	"github.com/fdurand/moby/pkg/sysinfo"
)

const platformSupported = false

func setupResolvConf(config *config.Config) {
}

func getSysInfo(daemon *Daemon) *sysinfo.SysInfo {
	return sysinfo.New()
}
