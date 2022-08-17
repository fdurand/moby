package types // import "github.com/fdurand/moby/api/types"

import "github.com/fdurand/moby/api/types/volume"

// Volume volume
//
// Deprecated: use github.com/fdurand/moby/api/types/volume.Volume
type Volume = volume.Volume

// VolumeUsageData Usage details about the volume. This information is used by the
// `GET /system/df` endpoint, and omitted in other endpoints.
//
// Deprecated: use github.com/fdurand/moby/api/types/volume.UsageData
type VolumeUsageData = volume.UsageData
