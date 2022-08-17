package layer // import "github.com/fdurand/moby/layer"

import "github.com/docker/distribution"

var _ distribution.Describable = &roLayer{}

func (rl *roLayer) Descriptor() distribution.Descriptor {
	return rl.descriptor
}
