package config

import (
	"testing"

	"github.com/fdurand/moby/api/types/filters"
	"github.com/google/go-cmp/cmp"
	"gotest.tools/v3/assert"
	"gotest.tools/v3/fs"
)

func TestBuilderGC(t *testing.T) {
	tempFile := fs.NewFile(t, "config", fs.WithContent(`{
  "builder": {
    "gc": {
      "enabled": true,
      "policy": [
        {"keepStorage": "10GB", "filter": ["unused-for=2200h"]},
        {"keepStorage": "50GB", "filter": {"unused-for": {"3300h": true}}},
        {"keepStorage": "100GB", "all": true}
      ]
    }
  }
}`))
	defer tempFile.Remove()
	configFile := tempFile.Path()

	cfg, err := MergeDaemonConfigurations(&Config{}, nil, configFile)
	assert.NilError(t, err)
	assert.Assert(t, cfg.Builder.GC.Enabled)
	f1 := filters.NewArgs()
	f1.Add("unused-for", "2200h")
	f2 := filters.NewArgs()
	f2.Add("unused-for", "3300h")
	expectedPolicy := []BuilderGCRule{
		{KeepStorage: "10GB", Filter: BuilderGCFilter(f1)},
		{KeepStorage: "50GB", Filter: BuilderGCFilter(f2)}, /* parsed from deprecated form */
		{KeepStorage: "100GB", All: true},
	}
	assert.DeepEqual(t, cfg.Builder.GC.Policy, expectedPolicy, cmp.AllowUnexported(BuilderGCFilter{}))
	// double check to please the skeptics
	assert.Assert(t, filters.Args(cfg.Builder.GC.Policy[0].Filter).UniqueExactMatch("unused-for", "2200h"))
	assert.Assert(t, filters.Args(cfg.Builder.GC.Policy[1].Filter).UniqueExactMatch("unused-for", "3300h"))
}
