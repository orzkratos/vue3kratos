package demo1x

import "github.com/yyle88/runpath"

func GetClientBPath() string {
	return runpath.PARENT.Join("greeter.client.ts-B.txt")
}

func GetClientAPath() string {
	return runpath.PARENT.Join("greeter.client.ts-A.txt")
}
