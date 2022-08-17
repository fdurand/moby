//go:build !linux
// +build !linux

package cluster // import "github.com/fdurand/moby/daemon/cluster"

import "net"

func (c *Cluster) resolveSystemAddr() (net.IP, error) {
	return c.resolveSystemAddrViaSubnetCheck()
}
