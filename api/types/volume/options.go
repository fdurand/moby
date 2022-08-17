package volume // import "github.com/fdurand/moby/api/types/volume"

import "github.com/fdurand/moby/api/types/filters"

// ListOptions holds parameters to list volumes.
type ListOptions struct {
	Filters filters.Args
}
