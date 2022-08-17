package daemon // import "github.com/fdurand/moby/daemon"

import (
	"testing"

	containertypes "github.com/fdurand/moby/api/types/container"
	"github.com/fdurand/moby/container"
	"github.com/fdurand/moby/daemon/config"
	"github.com/fdurand/moby/daemon/exec"
	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestGetInspectData(t *testing.T) {
	c := &container.Container{
		ID:           "inspect-me",
		HostConfig:   &containertypes.HostConfig{},
		State:        container.NewState(),
		ExecCommands: exec.NewStore(),
	}

	d := &Daemon{
		linkIndex:   newLinkIndex(),
		configStore: &config.Config{},
	}

	_, err := d.getInspectData(c)
	assert.Check(t, is.ErrorContains(err, ""))

	c.Dead = true
	_, err = d.getInspectData(c)
	assert.Check(t, err)
}
