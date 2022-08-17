package containerd

import (
	"context"

	"github.com/fdurand/moby/api/types"
	"github.com/fdurand/moby/api/types/filters"
)

// ImagesPrune removes unused images
func (i *ImageService) ImagesPrune(ctx context.Context, pruneFilters filters.Args) (*types.ImagesPruneReport, error) {
	panic("not implemented")
}
