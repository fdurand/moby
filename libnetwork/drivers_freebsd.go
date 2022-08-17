package libnetwork

import (
	"github.com/fdurand/moby/libnetwork/drivers/null"
	"github.com/fdurand/moby/libnetwork/drivers/remote"
)

func getInitializers(experimental bool) []initializer {
	return []initializer{
		{null.Init, "null"},
		{remote.Init, "remote"},
	}
}
