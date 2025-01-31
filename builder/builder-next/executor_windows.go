package buildkit

import (
	"context"
	"errors"

	"github.com/fdurand/moby/daemon/config"
	"github.com/fdurand/moby/libnetwork"
	"github.com/fdurand/moby/pkg/idtools"
	"github.com/moby/buildkit/executor"
	"github.com/moby/buildkit/executor/oci"
)

func newExecutor(_, _ string, _ libnetwork.NetworkController, _ *oci.DNSConfig, _ bool, _ idtools.IdentityMapping, _ string) (executor.Executor, error) {
	return &winExecutor{}, nil
}

type winExecutor struct {
}

func (w *winExecutor) Run(ctx context.Context, id string, root executor.Mount, mounts []executor.Mount, process executor.ProcessInfo, started chan<- struct{}) (err error) {
	return errors.New("buildkit executor not implemented for windows")
}

func (w *winExecutor) Exec(ctx context.Context, id string, process executor.ProcessInfo) error {
	return errors.New("buildkit executor not implemented for windows")
}

func getDNSConfig(config.DNSConfig) *oci.DNSConfig {
	return nil
}
