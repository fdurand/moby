package containerd

import (
	"github.com/fdurand/moby/builder"
)

// MakeImageCache creates a stateful image cache.
func (i *ImageService) MakeImageCache(cacheFrom []string) builder.ImageCache {
	panic("not implemented")
}
