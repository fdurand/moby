package ioutils // import "github.com/fdurand/moby/pkg/ioutils"

import (
	"os"

	"github.com/fdurand/moby/pkg/longpath"
)

// TempDir is the equivalent of os.MkdirTemp, except that the result is in Windows longpath format.
func TempDir(dir, prefix string) (string, error) {
	tempDir, err := os.MkdirTemp(dir, prefix)
	if err != nil {
		return "", err
	}
	return longpath.AddPrefix(tempDir), nil
}
