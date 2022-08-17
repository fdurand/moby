package libnetwork

import (
	"github.com/fdurand/moby/libnetwork/drivers/bridge"
	"github.com/fdurand/moby/libnetwork/drivers/host"
	"github.com/fdurand/moby/libnetwork/drivers/ipvlan"
	"github.com/fdurand/moby/libnetwork/drivers/macvlan"
	"github.com/fdurand/moby/libnetwork/drivers/null"
	"github.com/fdurand/moby/libnetwork/drivers/overlay"
	"github.com/fdurand/moby/libnetwork/drivers/remote"
)

func getInitializers(experimental bool) []initializer {
	in := []initializer{
		{bridge.Init, "bridge"},
		{host.Init, "host"},
		{ipvlan.Init, "ipvlan"},
		{macvlan.Init, "macvlan"},
		{null.Init, "null"},
		{overlay.Init, "overlay"},
		{remote.Init, "remote"},
	}
	return in
}
