package libnetwork

import (
	"github.com/fdurand/moby/libnetwork/drvregistry"
	"github.com/fdurand/moby/libnetwork/ipamapi"
	builtinIpam "github.com/fdurand/moby/libnetwork/ipams/builtin"
	nullIpam "github.com/fdurand/moby/libnetwork/ipams/null"
	remoteIpam "github.com/fdurand/moby/libnetwork/ipams/remote"
	"github.com/fdurand/moby/libnetwork/ipamutils"
)

func initIPAMDrivers(r *drvregistry.DrvRegistry, lDs, gDs interface{}, addressPool []*ipamutils.NetworkToSplit) error {
	builtinIpam.SetDefaultIPAddressPool(addressPool)
	for _, fn := range [](func(ipamapi.Callback, interface{}, interface{}) error){
		builtinIpam.Init,
		remoteIpam.Init,
		nullIpam.Init,
	} {
		if err := fn(r, lDs, gDs); err != nil {
			return err
		}
	}

	return nil
}
