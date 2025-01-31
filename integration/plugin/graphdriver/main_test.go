package graphdriver // import "github.com/fdurand/moby/integration/plugin/graphdriver"

import (
	"fmt"
	"os"
	"testing"

	"github.com/fdurand/moby/pkg/reexec"
	"github.com/fdurand/moby/testutil/environment"
)

var (
	testEnv *environment.Execution
)

func init() {
	reexec.Init() // This is required for external graphdriver tests
}

func TestMain(m *testing.M) {
	var err error
	testEnv, err = environment.New()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = environment.EnsureFrozenImagesLinux(testEnv)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	testEnv.Print()
	os.Exit(m.Run())
}
