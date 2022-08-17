//go:build windows
// +build windows

package daemon // import "github.com/fdurand/moby/daemon"

import "github.com/fdurand/moby/pkg/plugingetter"

func registerMetricsPluginCallback(getter plugingetter.PluginGetter, sockPath string) {
}

func (daemon *Daemon) listenMetricsSock() (string, error) {
	return "", nil
}
