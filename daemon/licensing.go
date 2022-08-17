package daemon // import "github.com/fdurand/moby/daemon"

import (
	"github.com/fdurand/moby/api/types"
	"github.com/fdurand/moby/dockerversion"
)

func (daemon *Daemon) fillLicense(v *types.Info) {
	v.ProductLicense = dockerversion.DefaultProductLicense
}
