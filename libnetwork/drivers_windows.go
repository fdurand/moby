package libnetwork

import (
	"github.com/fdurand/moby/libnetwork/drivers/null"
	"github.com/fdurand/moby/libnetwork/drivers/remote"
	"github.com/fdurand/moby/libnetwork/drivers/windows"
	"github.com/fdurand/moby/libnetwork/drivers/windows/overlay"
)

func getInitializers(experimental bool) []initializer {
	return []initializer{
		{null.Init, "null"},
		{overlay.Init, "overlay"},
		{remote.Init, "remote"},
		{windows.GetInit("transparent"), "transparent"},
		{windows.GetInit("l2bridge"), "l2bridge"},
		{windows.GetInit("l2tunnel"), "l2tunnel"},
		{windows.GetInit("nat"), "nat"},
		{windows.GetInit("internal"), "internal"},
		{windows.GetInit("private"), "private"},
		{windows.GetInit("ics"), "ics"},
	}
}
