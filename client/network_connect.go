package client // import "github.com/fdurand/moby/client"

import (
	"context"

	"github.com/fdurand/moby/api/types"
	"github.com/fdurand/moby/api/types/network"
)

// NetworkConnect connects a container to an existent network in the docker host.
func (cli *Client) NetworkConnect(ctx context.Context, networkID, containerID string, config *network.EndpointSettings) error {
	nc := types.NetworkConnect{
		Container:      containerID,
		EndpointConfig: config,
	}
	resp, err := cli.post(ctx, "/networks/"+networkID+"/connect", nil, nc, nil)
	ensureReaderClosed(resp)
	return err
}
