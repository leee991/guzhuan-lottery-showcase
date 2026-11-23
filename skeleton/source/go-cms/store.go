// 接口定义 — source/go-cms/store.go
// 仅保留数据结构与函数签名。


package main


import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"sync"
	"time"
	"unicode/utf8"
)


type Content struct {
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
	Version   int    `json:"version"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	Author    string `json:"author"`
}

type Revision struct {
	Version   int     `json:"version"`
	CreatedAt string  `json:"createdAt"`
	Author    string  `json:"author"`
	Content   Content `json:"content"`
}

type Audit struct {
	ID       string `json:"id"`
	Operator string `json:"operator"`
	Action   string `json:"action"`
	Target   string `json:"target"`
	Time     string `json:"time"`
}

type database struct {
	SchemaVersion int                   `json:"schemaVersion"`
	Items         map[string]Content    `json:"items"`
	Revisions     map[string][]Revision `json:"revisions"`
	Audit         []Audit               `json:"audit"`
}

type Store struct {
	mu         sync.RWMutex
	dir        string
	db         database
	generation uint64
}

type snapshot struct {
	SchemaVersion int       `json:"schemaVersion"`
	CreatedAt     string    `json:"createdAt"`
	Items         []Content `json:"items"`
}

type backupFile struct {
	ID        string          `json:"id"`
	CreatedAt string          `json:"createdAt"`
	SHA256    string          `json:"sha256"`
	Payload   json.RawMessage `json:"payload"`
}

type BackupInfo struct {
	ID        string `json:"id"`
	CreatedAt string `json:"createdAt"`
	SHA256    string `json:"sha256"`
	SizeBytes int64  `json:"sizeBytes"`
}


func newID() string {
	// 内部业务实现略
	panic("not implemented")
}

func stamp() string          {
	// 内部业务实现略
	panic("not implemented")
}

func digest(b []byte) string {
	// 内部业务实现略
	panic("not implemented")
}


func NewStore(dir string) (*Store, error) {
	// 内部业务实现略
	panic("not implemented")
}

func atomicFile(path string, b []byte) error {
	// 内部业务实现略
	panic("not implemented")
}

func (s *Store) save(db database) error {
	// 内部业务实现略
	panic("not implemented")
}

func cloneDB(db database) database {
	// 内部业务实现略
	panic("not implemented")
}

func addAudit(db *database, actor, action, target string) {
	// 内部业务实现略
	panic("not implemented")
}

func addRevision(db *database, c Content) {
	// 内部业务实现略
	panic("not implemented")
}


func safeURL(raw string) bool {
	// 内部业务实现略
	panic("not implemented")
}

func validate(c Content) error {
	// 内部业务实现略
	panic("not implemented")
}

func unique(db database, c Content) bool {
	// 内部业务实现略
	panic("not implemented")
}

func (s *Store) put(c Content, actor string, create bool) (Content, error) {
	// 内部业务实现略
	panic("not implemented")
}

func sortedItems(db database) []Content {
	// 内部业务实现略
	panic("not implemented")
}

func (s *Store) restoreRevision(id string, revision, version int, actor string) (Content, error) {
	// 内部业务实现略
	panic("not implemented")
}

func (s *Store) backupListLocked() ([]BackupInfo, error) {
	// 内部业务实现略
	panic("not implemented")
}

func (s *Store) makeBackupLocked() (BackupInfo, error) {
	// 内部业务实现略
	panic("not implemented")
}

func (s *Store) backup(actor string) (BackupInfo, error) {
	// 内部业务实现略
	panic("not implemented")
}

func (s *Store) loadBackupLocked(id string) (backupFile, snapshot, error) {
	// 内部业务实现略
	panic("not implemented")
}

func (s *Store) verifyBackup(id, actor string) (backupFile, snapshot, error) {
	// 内部业务实现略
	panic("not implemented")
}


func (s *Store) restoreBackup(id, actor string) (string, int, error) {
	// 内部业务实现略
	panic("not implemented")
}
