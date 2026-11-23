// 接口定义 — source/history.go
// 仅保留数据结构与函数签名。

package main


import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log"
	"math"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
	"sync"
	"time"
)


type routeTrend struct {
	Route      string   `json:"route"`
	Requests   uint64   `json:"requests"`
	Errors     uint64   `json:"errors"`
	ErrorRate  *float64 `json:"errorRate"`
	MeanMillis *float64 `json:"meanMillis"`
	P95Millis  *float64 `json:"p95Millis"`
}

type performanceSample struct {
	At              string       `json:"at"`
	IntervalSeconds float64      `json:"intervalSeconds"`
	Requests        uint64       `json:"requests"`
	Errors          uint64       `json:"errors"`
	ErrorRate       *float64     `json:"errorRate"`
	MeanMillis      *float64     `json:"meanMillis"`
	P95Millis       *float64     `json:"p95Millis"`
	HeapAllocBytes  uint64       `json:"heapAllocBytes"`
	Goroutines      int          `json:"goroutines"`
	Host            hostSample   `json:"host"`
	Routes          []routeTrend `json:"routes"`
}

type metricBucket struct {
	requests, errors uint64
	total            float64
	durations        []float64
}


type historyFile struct {
	SchemaVersion int                 `json:"schemaVersion"`
	PersistedAt   string              `json:"persistedAt"`
	Samples       []performanceSample `json:"samples"`
}

type performanceHistory struct {
	mu               sync.Mutex
	sampleMu         sync.Mutex
	dir              string
	samples          []performanceSample
	window           metricBucket
	routes           map[string]*metricBucket
	intervalStart    time.Time
	previousCPU      cpuCounters
	persistedAt      string
	persistenceError bool
}


func (b *metricBucket) record(status int, ms float64) {
	// 内部业务实现略
	panic("not implemented")
}

func (b metricBucket) trend(route string) routeTrend {
	// 内部业务实现略
	panic("not implemented")
}

func normalizedRoute(method, path string) string {
	// 内部业务实现略
	panic("not implemented")
}

func validRouteTemplate(r string) bool {
	// 内部业务实现略
	panic("not implemented")
}

func newPerformanceHistory(dir string, now time.Time) (*performanceHistory, error) {
	// 内部业务实现略
	panic("not implemented")
}

func (h *performanceHistory) record(route string, status int, ms float64) {
	// 内部业务实现略
	panic("not implemented")
}

func topRoutes(routes []routeTrend) []routeTrend {
	// 内部业务实现略
	panic("not implemented")
}

func aggregateRoutes(routes []routeTrend) []routeTrend {
	// 内部业务实现略
	panic("not implemented")
}

func (h *performanceHistory) collect(now time.Time) error {
	// 内部业务实现略
	panic("not implemented")
}

func (s *Server) sampleHistory(ctx context.Context) {
	// 内部业务实现略
	panic("not implemented")
}

func (s *Server) historyRoute(w http.ResponseWriter, r *http.Request) {
	// 内部业务实现略
	panic("not implemented")
}
