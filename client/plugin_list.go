package client // import "github.com/fdurand/moby/client"

import (
	"context"
	"encoding/json"
	"net/url"

	"github.com/fdurand/moby/api/types"
	"github.com/fdurand/moby/api/types/filters"
)

// PluginList returns the installed plugins
func (cli *Client) PluginList(ctx context.Context, filter filters.Args) (types.PluginsListResponse, error) {
	var plugins types.PluginsListResponse
	query := url.Values{}

	if filter.Len() > 0 {
		//nolint:staticcheck // ignore SA1019 for old code
		filterJSON, err := filters.ToParamWithVersion(cli.version, filter)
		if err != nil {
			return plugins, err
		}
		query.Set("filters", filterJSON)
	}
	resp, err := cli.get(ctx, "/plugins", query, nil)
	defer ensureReaderClosed(resp)
	if err != nil {
		return plugins, err
	}

	err = json.NewDecoder(resp.body).Decode(&plugins)
	return plugins, err
}
