package daemon // import "github.com/fdurand/moby/daemon"

import (
	"testing"

	"github.com/fdurand/moby/api/types"
	"github.com/fdurand/moby/dockerversion"
	"gotest.tools/v3/assert"
)

func TestFillLicense(t *testing.T) {
	v := &types.Info{}
	d := &Daemon{
		root: "/var/lib/docker/",
	}
	d.fillLicense(v)
	assert.Assert(t, v.ProductLicense == dockerversion.DefaultProductLicense)
}
