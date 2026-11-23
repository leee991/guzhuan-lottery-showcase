// 接口定义 — source/go-cms/server_test.go
// 仅保留数据结构与函数签名。


package main


import (
	"bytes"
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


func setup(t *testing.T) (*Server, string) {
	// 内部业务实现略
	panic("not implemented")
}

func call(s *Server, method, path, role string, body any) *httptest.ResponseRecorder {
	// 内部业务实现略
	panic("not implemented")
}

func assertStatus(t *testing.T, w *httptest.ResponseRecorder, want int) {
	// 内部业务实现略
	panic("not implemented")
}

func content(slug string) Content {
	// 内部业务实现略
	panic("not implemented")
}

func create(t *testing.T, s *Server, c Content) Content {
	// 内部业务实现略
	panic("not implemented")
}

func publicItems(t *testing.T, w *httptest.ResponseRecorder) []map[string]any {
	// 内部业务实现略
	panic("not implemented")
}


func TestAuthAndRoleBoundaries(t *testing.T) {
	// 内部业务实现略
	panic("not implemented")
}

func TestAuthFailsClosedOnRedirectOrMalformedResponse(t *testing.T) {
	// 内部业务实现略
	panic("not implemented")
}

func TestDraftPrivateScheduledPublicationAndETag(t *testing.T) {
	// 内部业务实现略
	panic("not implemented")
}

func TestValidationLimitsAndSafeURLs(t *testing.T) {
	// 内部业务实现略
	panic("not implemented")
}

func TestConflictRevisionsAuditPersistenceAndCacheInvalidation(t *testing.T) {
	// 内部业务实现略
	panic("not implemented")
}

func TestConcurrentVersionAndSlugProtection(t *testing.T) {
	// 内部业务实现略
	panic("not implemented")
}

func TestBackupVerifyRestoreSafetyAndTamper(t *testing.T) {
	// 内部业务实现略
	panic("not implemented")
}

func TestBackupRetentionAndReadiness(t *testing.T) {
	// 内部业务实现略
	panic("not implemented")
}

func TestOversizedSaveDoesNotChangePersistentStore(t *testing.T) {
	// 内部业务实现略
	panic("not implemented")
}


func TestMetricsAreRealAndRevisionsBounded(t *testing.T) {
	// 内部业务实现略
	panic("not implemented")
}
