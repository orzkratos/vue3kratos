package demo1x

import (
	"github.com/yyle88/osexistpath/osmustexist"
	"github.com/yyle88/runpath"
)

func GetProjectPath() string {
	return osmustexist.ROOT(runpath.PARENT.Join("demo1go"))
}

func GetClientBPath() string {
	return runpath.PARENT.Join("greeter.client.ts-B.txt")
}

func GetClientAPath() string {
	return runpath.PARENT.Join("greeter.client.ts-A.txt")
}
