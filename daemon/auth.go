package daemon // import "github.com/fdurand/moby/daemon"

import (
	"context"

	"github.com/fdurand/moby/api/types/registry"
	"github.com/fdurand/moby/dockerversion"
)

// AuthenticateToRegistry checks the validity of credentials in authConfig
func (daemon *Daemon) AuthenticateToRegistry(ctx context.Context, authConfig *registry.AuthConfig) (string, string, error) {
	return daemon.registryService.Auth(ctx, authConfig, dockerversion.DockerUserAgent(ctx))
}
