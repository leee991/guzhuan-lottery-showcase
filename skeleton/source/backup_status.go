// 接口定义 — source/backup_status.go
// 仅保留数据结构与函数签名。


package main


import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"
)


type backupStatus struct {
	Configured    bool   `json:"configured"`
	Status        string `json:"status"`
	StartedAt     string `json:"startedAt"`
	FinishedAt    string `json:"finishedAt"`
	LastSuccessAt string `json:"lastSuccessAt"`
	SnapshotID    string `json:"snapshotId"`
	Files         uint64 `json:"files"`
	TotalBytes    uint64 `json:"totalBytes"`
	Retention     struct {
		Daily  int `json:"daily"`
		Weekly int `json:"weekly"`
	} `json:"retention"`
	Restore struct {
		Status       string `json:"status"`
		At           string `json:"at"`
		CheckedFiles uint64 `json:"checkedFiles"`
		Message      string `json:"message"`
	} `json:"restore"`
	Message string `json:"message"`
}


func readBackupStatus() (backupStatus, error) {
	// 内部业务实现略
	panic("not implemented")
}


func backupStatusRoute(w http.ResponseWriter) {
	// 内部业务实现略
	panic("not implemented")
}

func backupDegraded(now time.Time) bool {
	// 内部业务实现略
	panic("not implemented")
}


func backupRequired() bool {
	// 内部业务实现略
	panic("not implemented")
}
