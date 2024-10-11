package utils

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/runpath"
)

func TestWalkFiles(t *testing.T) {
	require.NoError(t, WalkFiles(runpath.PARENT.Path(), func(path string, info os.FileInfo) error {
		t.Log(path)
		return nil
	}))
}
