// 接口定义 — source/go-cms/server.go
// 仅保留数据结构与函数签名。


package main


import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
	"sync"
	"time"
)


type Actor struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	Status   string `json:"status,omitempty"`
}

type publicContent struct {
	ID        string `json:"id"`
	Kind      string `json:"kind"`
	Title     string `json:"title"`
	Slug      string `json:"slug"`
	Category  string `json:"category"`
	Summary   string `json:"summary"`
	Body      string `json:"body"`
	CoverURL  string `json:"coverUrl"`
	LinkURL   string `json:"linkUrl"`
	Status    string `json:"status"`
	PublishAt string `json:"publishAt"`
	Sort      int    `json:"sort"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type cacheEntry struct {
	generation uint64
	expires    time.Time
	etag       string
	body       []byte
}

type metrics struct {
	mu        sync.Mutex
	requests  uint64
	errors    uint64
	cacheHits uint64
	durations []float64
}

type Server struct {
	store   *Store
	authURL string
	client  *http.Client
	started time.Time
	metrics metrics
	cacheMu sync.Mutex
	cache   map[string]cacheEntry
}

type responseWriter struct {
	http.ResponseWriter
	status int
}


func (w *responseWriter) WriteHeader(n int) {
	// 内部业务实现略
	panic("not implemented")
}

func (w *responseWriter) Write(p []byte) (int, error) {
	// 内部业务实现略
	panic("not implemented")
}


func NewServer(dir, authURL string) (*Server, error) {
	// 内部业务实现略
	panic("not implemented")
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 内部业务实现略
	panic("not implemented")
}

func respond(w http.ResponseWriter, status int, v any) {
	// 内部业务实现略
	panic("not implemented")
}

func failure(w http.ResponseWriter, status int, code, msg string) {
	// 内部业务实现略
	panic("not implemented")
}

func method(w http.ResponseWriter, r *http.Request, allowed string) bool {
	// 内部业务实现略
	panic("not implemented")
}

func decode(w http.ResponseWriter, r *http.Request, v any) bool {
	// 内部业务实现略
	panic("not implemented")
}

func writeAllowed(a Actor) bool {
	// 内部业务实现略
	panic("not implemented")
}

func recognized(role string) bool {
	// 内部业务实现略
	panic("not implemented")
}

func (s *Server) actor(r *http.Request) (Actor, error) {
	// 内部业务实现略
	panic("not implemented")
}

func storageError(w http.ResponseWriter, err error) {
	// 内部业务实现略
	panic("not implemented")
}

func (s *Server) route(w http.ResponseWriter, r *http.Request) {
	// 内部业务实现略
	panic("not implemented")
}

func (s *Server) contentRoute(w http.ResponseWriter, r *http.Request, a Actor, tail string) {
	// 内部业务实现略
	panic("not implemented")
}

func (s *Server) backupRoute(w http.ResponseWriter, r *http.Request, a Actor, tail string) {
	// 内部业务实现略
	panic("not implemented")
}

func (s *Server) public(w http.ResponseWriter, r *http.Request) {
	// 内部业务实现略
	panic("not implemented")
}

func writePublic(w http.ResponseWriter, r *http.Request, c cacheEntry) {
	// 内部业务实现略
	panic("not implemented")
}

func (s *Server) overview(w http.ResponseWriter) {
	// 内部业务实现略
	panic("not implemented")
}

func (s *Server) autoBackups(ctx context.Context) {
	// 内部业务实现略
	panic("not implemented")
}


func (s *Server) ready() error {
	// 内部业务实现略
	panic("not implemented")
}
