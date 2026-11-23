// 接口定义 — source/editorial_test.go
// 仅保留数据结构与函数签名。


package main


import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)


func taxonFrom(t *testing.T, w *httptest.ResponseRecorder, status int) Taxon {
	// 内部业务实现略
	panic("not implemented")
}

func createTaxon(t *testing.T, s *Server, kind, name, parent string) Taxon {
	// 内部业务实现略
	panic("not implemented")
}

func draftFrom(t *testing.T, w *httptest.ResponseRecorder) *RecoveryDraft {
	// 内部业务实现略
	panic("not implemented")
}

func TestTaxonomyCRUDVersionsAndReferences(t *testing.T) {
	// 内部业务实现略
	panic("not implemented")
}

func TestTaxonomyCyclesDepthAndAtomicFailures(t *testing.T) {
	// 内部业务实现略
	panic("not implemented")
}

func actors(t *testing.T, s *Server) {
	// 内部业务实现略
	panic("not implemented")
}

func TestPrivateRecoveryIsolationTTLConflictAndRestart(t *testing.T) {
	// 内部业务实现略
	panic("not implemented")
}

func TestPrivateRecoveryLimitsConcurrentAndQuota(t *testing.T) {
	// 内部业务实现略
	panic("not implemented")
}

func TestEditorialCorruptionAndPendingAudit(t *testing.T) {
	// 内部业务实现略
	panic("not implemented")
}
