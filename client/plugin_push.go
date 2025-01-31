package client // import "github.com/fdurand/moby/client"

import (
	"context"
	"io"

	"github.com/fdurand/moby/api/types/registry"
)

// PluginPush pushes a plugin to a registry
func (cli *Client) PluginPush(ctx context.Context, name string, registryAuth string) (io.ReadCloser, error) {
	headers := map[string][]string{registry.AuthHeader: {registryAuth}}
	resp, err := cli.post(ctx, "/plugins/"+name+"/push", nil, nil, headers)
	if err != nil {
		return nil, err
	}
	return resp.body, nil
}
