package client // import "github.com/fdurand/moby/client"

import (
	"context"
	"encoding/json"
	"net/url"

	"github.com/fdurand/moby/api/types"
	"github.com/fdurand/moby/api/types/filters"
	"github.com/fdurand/moby/api/types/swarm"
)

// NodeList returns the list of nodes.
func (cli *Client) NodeList(ctx context.Context, options types.NodeListOptions) ([]swarm.Node, error) {
	query := url.Values{}

	if options.Filters.Len() > 0 {
		filterJSON, err := filters.ToJSON(options.Filters)

		if err != nil {
			return nil, err
		}

		query.Set("filters", filterJSON)
	}

	resp, err := cli.get(ctx, "/nodes", query, nil)
	defer ensureReaderClosed(resp)
	if err != nil {
		return nil, err
	}

	var nodes []swarm.Node
	err = json.NewDecoder(resp.body).Decode(&nodes)
	return nodes, err
}
