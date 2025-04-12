package utils

import (
	"encoding/hex"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/yyle88/erero"
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
