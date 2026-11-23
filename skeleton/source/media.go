// 接口定义 — source/media.go
// 仅保留数据结构与函数签名。

package main


import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"io"
	"log"
	"math"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"
)

type Media struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	URL       string `json:"url"`
	MIME      string `json:"mime"`
	Width     int    `json:"width"`
	Height    int    `json:"height"`
	SizeBytes int64  `json:"sizeBytes"`
	CreatedAt string `json:"createdAt"`
	Author    string `json:"author"`
	Archived  bool   `json:"archived"`
}

type mediaDatabase struct {
	SchemaVersion int              `json:"schemaVersion"`
	Items         map[string]Media `json:"items"`
	PendingAudit  *Audit           `json:"pendingAudit,omitempty"`
}


type mediaError struct {
	status        int
	code, message string
}


func (e *mediaError) Error() string {
	// 内部业务实现略
	panic("not implemented")
}

func mediaFailure(w http.ResponseWriter, err error) {
	// 内部业务实现略
	panic("not implemented")
}

func invalidMedia(msg string) error {
	// 内部业务实现略
	panic("not implemented")
}

func mediaExtension(mime string) string {
	// 内部业务实现略
	panic("not implemented")
}

func mediaURL(id, mime string) string {
	// 内部业务实现略
	panic("not implemented")
}

func validMediaName(name string) bool {
	// 内部业务实现略
	panic("not implemented")
}

func cloneMedia(d mediaDatabase) mediaDatabase {
	// 内部业务实现略
	panic("not implemented")
}

func (s *Store) mediaPath(m Media) string {
	// 内部业务实现略
	panic("not implemented")
}

func (s *Store) saveMedia(d mediaDatabase) error {
	// 内部业务实现略
	panic("not implemented")
}

func atomicMediaFile(path string, b []byte) error {
	// 内部业务实现略
	panic("not implemented")
}

func atomicMediaFileWithWriter(path string, b []byte, write func(string, []byte) error) error {
	// 内部业务实现略
	panic("not implemented")
}

func regularPrivate(path string) (os.FileInfo, error) {
	// 内部业务实现略
	panic("not implemented")
}

func (s *Store) initMedia() error {
	// 内部业务实现略
	panic("not implemented")
}

func (s *Store) mediaUsageLocked() (int, int64, error) {
	// 内部业务实现略
	panic("not implemented")
}

func (s *Store) mediaReadyLocked() error {
	// 内部业务实现略
	panic("not implemented")
}

func mergedMediaAudit(items []Audit, pending *Audit) []Audit {
	// 内部业务实现略
	panic("not implemented")
}

func (s *Store) flushMediaAuditLocked() error {
	// 内部业务实现略
	panic("not implemented")
}

func (s *Store) commitMediaLocked(n mediaDatabase, actor, action, target string) error {
	// 内部业务实现略
	panic("not implemented")
}

func (s *Store) putMedia(name, mime string, width, height int, data []byte, actor string) (Media, error) {
	// 内部业务实现略
	panic("not implemented")
}

func (s *Store) archiveMedia(id string, archived bool, actor string) (Media, error) {
	// 内部业务实现略
	panic("not implemented")
}


func readMediaUpload(w http.ResponseWriter, r *http.Request) ([]byte, string, error) {
	// 内部业务实现略
	panic("not implemented")
}

func uploadReadError(err error) error {
	// 内部业务实现略
	panic("not implemented")
}

func resizeMedia(src image.Image) image.Image {
	// 内部业务实现略
	panic("not implemented")
}

func normalizeMedia(data []byte) ([]byte, string, int, int, error) {
	// 内部业务实现略
	panic("not implemented")
}

func (s *Server) mediaRoute(w http.ResponseWriter, r *http.Request, a Actor) {
	// 内部业务实现略
	panic("not implemented")
}

func (s *Server) publicMedia(w http.ResponseWriter, r *http.Request) {
	// 内部业务实现略
	panic("not implemented")
}
