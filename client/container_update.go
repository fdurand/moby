package client // import "github.com/fdurand/moby/client"

import (
	"context"
	"encoding/json"

	"github.com/fdurand/moby/api/types/container"
)

// ContainerUpdate updates the resources of a container.
func (cli *Client) ContainerUpdate(ctx context.Context, containerID string, updateConfig container.UpdateConfig) (container.ContainerUpdateOKBody, error) {
	var response container.ContainerUpdateOKBody
	serverResp, err := cli.post(ctx, "/containers/"+containerID+"/update", nil, updateConfig, nil)
	defer ensureReaderClosed(serverResp)
	if err != nil {
		return response, err
	}

	err = json.NewDecoder(serverResp.body).Decode(&response)
	return response, err
}
