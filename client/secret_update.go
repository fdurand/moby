package client // import "github.com/fdurand/moby/client"

import (
	"context"
	"net/url"

	"github.com/fdurand/moby/api/types/swarm"
)

// SecretUpdate attempts to update a secret.
func (cli *Client) SecretUpdate(ctx context.Context, id string, version swarm.Version, secret swarm.SecretSpec) error {
	if err := cli.NewVersionError("1.25", "secret update"); err != nil {
		return err
	}
	query := url.Values{}
	query.Set("version", version.String())
	resp, err := cli.post(ctx, "/secrets/"+id+"/update", query, secret, nil)
	ensureReaderClosed(resp)
	return err
}
