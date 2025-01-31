package containerd

import (
	"github.com/containerd/containerd/remotes"
	"github.com/containerd/containerd/remotes/docker"
	registrytypes "github.com/fdurand/moby/api/types/registry"
	"github.com/fdurand/moby/registry"
	"github.com/sirupsen/logrus"
)

func newResolverFromAuthConfig(authConfig *registrytypes.AuthConfig) remotes.Resolver {
	opts := []docker.RegistryOpt{}
	if authConfig != nil {
		cfgHost := registry.ConvertToHostname(authConfig.ServerAddress)
		if cfgHost == registry.IndexHostname {
			cfgHost = registry.DefaultRegistryHost
		}
		authorizer := docker.NewDockerAuthorizer(docker.WithAuthCreds(func(host string) (string, string, error) {
			if cfgHost != host {
				logrus.WithField("host", host).WithField("cfgHost", cfgHost).Warn("Host doesn't match")
				return "", "", nil
			}
			if authConfig.IdentityToken != "" {
				return "", authConfig.IdentityToken, nil
			}
			return authConfig.Username, authConfig.Password, nil
		}))

		opts = append(opts, docker.WithAuthorizer(authorizer))
	}

	return docker.NewResolver(docker.ResolverOptions{
		Hosts: docker.ConfigureDefaultRegistries(opts...),
	})
}
