package daemon // import "github.com/fdurand/moby/daemon"

import (
	"testing"

	"github.com/fdurand/moby/api/types/network"
	"github.com/fdurand/moby/errdefs"
	"gotest.tools/v3/assert"
)

// Test case for 35752
func TestVerifyNetworkingConfig(t *testing.T) {
	name := "mynet"
	endpoints := make(map[string]*network.EndpointSettings, 1)
	endpoints[name] = nil
	nwConfig := &network.NetworkingConfig{
		EndpointsConfig: endpoints,
	}
	err := verifyNetworkingConfig(nwConfig)
	assert.Check(t, errdefs.IsInvalidParameter(err))
}
