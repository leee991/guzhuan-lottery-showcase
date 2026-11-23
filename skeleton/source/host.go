// 接口定义 — source/host.go
// 仅保留数据结构与函数签名。

package main


type hostSample struct {
	CPUPercent      *float64 `json:"cpuPercent"`
	RAMUsedBytes    *uint64  `json:"ramUsedBytes"`
	RAMTotalBytes   *uint64  `json:"ramTotalBytes"`
	DiskUsedBytes   *uint64  `json:"diskUsedBytes"`
	DiskTotalBytes  *uint64  `json:"diskTotalBytes"`
	DiskUsedPercent *float64 `json:"diskUsedPercent"`
}

type cpuCounters struct {
	total, idle uint64
	valid       bool
}


func diskCritical(path string) bool {
	// 内部业务实现略
	panic("not implemented")
}
