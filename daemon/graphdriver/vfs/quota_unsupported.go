//go:build !linux
// +build !linux

package vfs // import "github.com/fdurand/moby/daemon/graphdriver/vfs"

import "github.com/fdurand/moby/quota"

type driverQuota struct {
}

func setupDriverQuota(driver *Driver) error {
	return nil
}

func (d *Driver) setQuotaOpt(size uint64) error {
	return quota.ErrQuotaNotSupported
}

func (d *Driver) getQuotaOpt() uint64 {
	return 0
}

func (d *Driver) setupQuota(dir string, size uint64) error {
	return quota.ErrQuotaNotSupported
}

func (d *Driver) quotaSupported() bool {
	return false
}
