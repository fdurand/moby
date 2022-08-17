//go:build !windows
// +build !windows

package libnetwork

import "github.com/fdurand/moby/libnetwork/ipamapi"

// Stub implementations for DNS related functions

func (n *network) startResolver() {
}

func defaultIpamForNetworkType(networkType string) string {
	return ipamapi.DefaultIPAM
}
