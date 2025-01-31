package images // import "github.com/fdurand/moby/daemon/images"

import (
	"context"
	"io"
	"time"

	"github.com/docker/distribution/manifest/schema2"
	"github.com/docker/distribution/reference"
	"github.com/fdurand/moby/api/types/registry"
	"github.com/fdurand/moby/distribution"
	progressutils "github.com/fdurand/moby/distribution/utils"
	"github.com/fdurand/moby/pkg/progress"
)

// PushImage initiates a push operation on the repository named localName.
func (i *ImageService) PushImage(ctx context.Context, image, tag string, metaHeaders map[string][]string, authConfig *registry.AuthConfig, outStream io.Writer) error {
	start := time.Now()
	ref, err := reference.ParseNormalizedNamed(image)
	if err != nil {
		return err
	}
	if tag != "" {
		// Push by digest is not supported, so only tags are supported.
		ref, err = reference.WithTag(ref, tag)
		if err != nil {
			return err
		}
	}

	// Include a buffer so that slow client connections don't affect
	// transfer performance.
	progressChan := make(chan progress.Progress, 100)

	writesDone := make(chan struct{})

	ctx, cancelFunc := context.WithCancel(ctx)

	go func() {
		progressutils.WriteDistributionProgress(cancelFunc, outStream, progressChan)
		close(writesDone)
	}()

	imagePushConfig := &distribution.ImagePushConfig{
		Config: distribution.Config{
			MetaHeaders:      metaHeaders,
			AuthConfig:       authConfig,
			ProgressOutput:   progress.ChanOutput(progressChan),
			RegistryService:  i.registryService,
			ImageEventLogger: i.LogImageEvent,
			MetadataStore:    i.distributionMetadataStore,
			ImageStore:       distribution.NewImageConfigStoreFromStore(i.imageStore),
			ReferenceStore:   i.referenceStore,
		},
		ConfigMediaType: schema2.MediaTypeImageConfig,
		LayerStores:     distribution.NewLayerProvidersFromStore(i.layerStore),
		TrustKey:        i.trustKey,
		UploadManager:   i.uploadManager,
	}

	err = distribution.Push(ctx, ref, imagePushConfig)
	close(progressChan)
	<-writesDone
	imageActions.WithValues("push").UpdateSince(start)
	return err
}
