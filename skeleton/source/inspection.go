// 接口定义 — source/inspection.go
// 仅保留数据结构与函数签名。

package main


import (
	"context"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"html"
	"image"
	_ "image/gif"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"
)


type inspectionIssue struct {
	Severity  string   `json:"severity"`
	Code      string   `json:"code"`
	Target    string   `json:"target"`
	Message   string   `json:"message"`
	Path      string   `json:"path,omitempty"`
	ContentID string   `json:"contentId,omitempty"`
	Title     string   `json:"title,omitempty"`
	Count     int      `json:"count,omitempty"`
	Examples  []string `json:"examples,omitempty"`
}

type inspectionScope struct {
	ContentCount        int    `json:"contentCount"`
	MediaCount          int    `json:"mediaCount"`
	LegacyPages         int    `json:"legacyPages"`
	StaticRootAvailable bool   `json:"staticRootAvailable"`
	ExternalFetch       bool   `json:"externalFetch"`
	Truncated           bool   `json:"truncated"`
	StopReason          string `json:"stopReason,omitempty"`
	AssetReferences     int    `json:"assetReferences"`
	Stylesheets         int    `json:"stylesheets"`
}

type inspectionReport struct {
	CheckedAt string            `json:"checkedAt"`
	Summary   map[string]int    `json:"summary"`
	Issues    []inspectionIssue `json:"issues"`
	Scope     inspectionScope   `json:"scope"`
}

type contentInspector struct {
	server     *Server
	root       string
	manifest   map[string]string
	media      map[string]Media
	report     inspectionReport
	seen       map[string]bool
	ctx        context.Context
	halted     bool
	matchCount int
	cssGroups  map[string]*inspectionIssue
}


func (i *contentInspector) halt(reason string) {
	// 内部业务实现略
	panic("not implemented")
}

func (i *contentInspector) stopped() bool {
	// 内部业务实现略
	panic("not implemented")
}

func (i *contentInspector) matches(pattern *regexp.Regexp, text string) [][]string {
	// 内部业务实现略
	panic("not implemented")
}


func (i *contentInspector) issue(severity, code, target, message, asset string, c *Content) {
	// 内部业务实现略
	panic("not implemented")
}

func (i *contentInspector) appendIssue(issue inspectionIssue) {
	// 内部业务实现略
	panic("not implemented")
}

func (i *contentInspector) cssIssue(severity, code, target, message, asset string) {
	// 内部业务实现略
	panic("not implemented")
}

func (i *contentInspector) flushCSSIssues() {
	// 内部业务实现略
	panic("not implemented")
}

func newInspector(s *Server, ctx context.Context) *contentInspector {
	// 内部业务实现略
	panic("not implemented")
}

func readLimitedFile(p string, max int64) ([]byte, error) {
	// 内部业务实现略
	panic("not implemented")
}

func (i *contentInspector) staticFile(publicPath string) (string, error) {
	// 内部业务实现略
	panic("not implemented")
}

func (i *contentInspector) localPath(raw string) (string, string) {
	// 内部业务实现略
	panic("not implemented")
}

func (i *contentInspector) asset(raw, target string, imageExpected bool, c *Content, mode string) {
	// 内部业务实现略
	panic("not implemented")
}

func imageDimensions(p string) (int, int, error) {
	// 内部业务实现略
	panic("not implemented")
}

func (i *contentInspector) content(c Content, all []Content) {
	// 内部业务实现略
	panic("not implemented")
}


func resolveLegacy(raw, base string) string {
	// 内部业务实现略
	panic("not implemented")
}

func (i *contentInspector) legacy() {
	// 内部业务实现略
	panic("not implemented")
}

func (s *Server) inspectionRoute(w http.ResponseWriter, r *http.Request) {
	// 内部业务实现略
	panic("not implemented")
}
