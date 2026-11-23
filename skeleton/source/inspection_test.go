// 接口定义 — source/inspection_test.go
// 仅保留数据结构与函数签名。

package main


import (
	"context"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"sync/atomic"
	"testing"
	"time"
)


func staticFixture(t *testing.T, root, p string, b []byte) {
	// 内部业务实现略
	panic("not implemented")
}

func reportBody(t *testing.T, w *httptest.ResponseRecorder) inspectionReport {
	// 内部业务实现略
	panic("not implemented")
}

func issueFound(r inspectionReport, code string) bool {
	// 内部业务实现略
	panic("not implemented")
}

func prepublish(s *Server, c Content) *httptest.ResponseRecorder {
	// 内部业务实现略
	panic("not implemented")
}


func TestInspectionContentChecksAndNoMutations(t *testing.T) {
	// 内部业务实现略
	panic("not implemented")
}

func TestInspectionNoFetchAndPathBoundaries(t *testing.T) {
	// 内部业务实现略
	panic("not implemented")
}

func TestInspectionMediaAndLegacyEvidence(t *testing.T) {
	// 内部业务实现略
	panic("not implemented")
}

func TestInspectionWebPDimensionsAndLimits(t *testing.T) {
	// 内部业务实现略
	panic("not implemented")
}

func TestBackupHealthDegradedWithoutDetails(t *testing.T) {
	// 内部业务实现略
	panic("not implemented")
}


func TestBackupRequiredConfiguration(t *testing.T) {
	// 内部业务实现略
	panic("not implemented")
}


func TestInspectionCancellationAndStrictLimits(t *testing.T) {
	// 内部业务实现略
	panic("not implemented")
}


func TestInspectionAssetAndCSSQueueLimits(t *testing.T) {
	// 内部业务实现略
	panic("not implemented")
}


func TestInspectionCSSAggregationPreservesHTMLFindings(t *testing.T) {
	// 内部业务实现略
	panic("not implemented")
}
