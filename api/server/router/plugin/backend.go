package plugin // import "github.com/fdurand/moby/api/server/router/plugin"

import (
	"context"
	"io"
	"net/http"

	"github.com/docker/distribution/reference"
	"github.com/fdurand/moby/api/types"
	"github.com/fdurand/moby/api/types/filters"
	"github.com/fdurand/moby/api/types/registry"
	"github.com/fdurand/moby/plugin"
)

// Backend for Plugin
type Backend interface {
	Disable(name string, config *types.PluginDisableConfig) error
	Enable(name string, config *types.PluginEnableConfig) error
	List(filters.Args) ([]types.Plugin, error)
	Inspect(name string) (*types.Plugin, error)
	Remove(name string, config *types.PluginRmConfig) error
	Set(name string, args []string) error
	Privileges(ctx context.Context, ref reference.Named, metaHeaders http.Header, authConfig *registry.AuthConfig) (types.PluginPrivileges, error)
	Pull(ctx context.Context, ref reference.Named, name string, metaHeaders http.Header, authConfig *registry.AuthConfig, privileges types.PluginPrivileges, outStream io.Writer, opts ...plugin.CreateOpt) error
	Push(ctx context.Context, name string, metaHeaders http.Header, authConfig *registry.AuthConfig, outStream io.Writer) error
	Upgrade(ctx context.Context, ref reference.Named, name string, metaHeaders http.Header, authConfig *registry.AuthConfig, privileges types.PluginPrivileges, outStream io.Writer) error
	CreateFromContext(ctx context.Context, tarCtx io.ReadCloser, options *types.PluginCreateOptions) error
}
