package utils

import (
	"encoding/hex"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
	"github.com/yyle88/erero"
	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
)

func NewUUID() string {
	uux := uuid.New()
	return hex.EncodeToString(uux[:])
}

func ExecInPath(path string, name string, args ...string) ([]byte, error) {
	if path == "" {
		return nil, erero.New("NO DIRECTORY PATH")
	}
	if strings.Contains(name, " ") {
		return nil, erero.New("COMMAND NAME CANNOT CONTAIN SPACES")
	}
	//这里只是打印出日志，以便于从日志中拷贝出命令直接手动执行
	zaplog.ZAPS.Skip1.LOG.Debug("EXEC_IN_PATH:", zap.String("cms", format1x(path, name, args...)))
	//设置命令和参数
	cmd := exec.Command(name, args...)
	//这是运行命令所在的目录
	cmd.Dir = path
	//执行命令且得到结果
	return cmd.CombinedOutput() //当错误时输出里有错误信息
}

func format1x(path string, name string, args ...string) string {
	return strings.TrimSpace(fmt.Sprintf("cd %s && %s", path, format2x(name, args...)))
}

func format2x(name string, args ...string) string {
	return strings.TrimSpace(fmt.Sprintf("%s %s", name, strings.Join(args, " ")))
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
