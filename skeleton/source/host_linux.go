// 接口定义 — source/host_linux.go
// 仅保留数据结构与函数签名。


package main


import (
	"os"
	"strconv"
	"strings"
	"syscall"
)


func hostDisk(path string) (uint64, uint64, bool) {
	// 内部业务实现略
	panic("not implemented")
}

func hostRead(path string, prev cpuCounters) (hostSample, cpuCounters) {
	// 内部业务实现略
	panic("not implemented")
}
