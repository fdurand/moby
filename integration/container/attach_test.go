package container // import "github.com/fdurand/moby/integration/container"

import (
	"context"
	"testing"

	"github.com/fdurand/moby/api/types"
	"github.com/fdurand/moby/api/types/container"
	"github.com/fdurand/moby/api/types/network"
	"gotest.tools/v3/assert"
)

func TestAttachWithTTY(t *testing.T) {
	testAttach(t, true, types.MediaTypeRawStream)
}

func TestAttachWithoutTTy(t *testing.T) {
	testAttach(t, false, types.MediaTypeMultiplexedStream)
}

func testAttach(t *testing.T, tty bool, expected string) {
	defer setupTest(t)()
	client := testEnv.APIClient()

	resp, err := client.ContainerCreate(context.Background(),
		&container.Config{
			Image: "busybox",
			Cmd:   []string{"echo", "hello"},
			Tty:   tty,
		},
		&container.HostConfig{},
		&network.NetworkingConfig{},
		nil,
		"",
	)
	assert.NilError(t, err)
	container := resp.ID
	defer client.ContainerRemove(context.Background(), container, types.ContainerRemoveOptions{
		Force: true,
	})

	attach, err := client.ContainerAttach(context.Background(), container, types.ContainerAttachOptions{
		Stdout: true,
		Stderr: true,
	})
	assert.NilError(t, err)
	mediaType, ok := attach.MediaType()
	assert.Check(t, ok)
	assert.Check(t, mediaType == expected)
}
