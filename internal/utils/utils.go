package utils

import (
	"encoding/hex"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/yyle88/erero"
	"github.com/yyle88/must"
	"github.com/yyle88/must/muststrings"
	"github.com/yyle88/osexistpath/osmustexist"
	"github.com/yyle88/rese"
	"github.com/yyle88/zaplog"
)

func NewUUID() string {
	newUUID := uuid.New()
	return hex.EncodeToString(newUUID[:])
}

func WalkFiles(root string, run func(path string, info os.FileInfo) error) error {
	if err := filepath.Walk(root,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return erero.Wro(err)
			}
			if info.IsDir() {
				return nil
			}
			return run(path, info)
		},
	); err != nil {
		return erero.Wro(err)
	}
	return nil
}

func CopyFiles(sourceRoot string, targetRoot string) {
	must.Done(WalkFiles(sourceRoot, func(srcPath string, info os.FileInfo) error {
		zaplog.SUG.Debugln(srcPath)
		muststrings.HasPrefix(srcPath, sourceRoot)

		relPath := must.V1(filepath.Rel(sourceRoot, srcPath))
		absPath := filepath.Join(sourceRoot, relPath)
		must.Same(absPath, srcPath)

		content := rese.V1(os.ReadFile(osmustexist.FILE(absPath)))
		dstPath := filepath.Join(targetRoot, relPath)
		zaplog.SUG.Debugln(dstPath)

		must.Done(os.MkdirAll(filepath.Dir(dstPath), 0755))
		must.Done(os.WriteFile(dstPath, content, 0644))
		return nil
	}))
}
