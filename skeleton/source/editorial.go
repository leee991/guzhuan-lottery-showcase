// 接口定义 — source/editorial.go
// 仅保留数据结构与函数签名。


package main


import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"
)


type Taxon struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	ParentID  string `json:"parentId"`
	Sort      int    `json:"sort"`
	Enabled   bool   `json:"enabled"`
	Version   int    `json:"version"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type taxonomyDatabase struct {
	SchemaVersion int              `json:"schemaVersion"`
	Categories    map[string]Taxon `json:"categories"`
	Tags          map[string]Taxon `json:"tags"`
	PendingAudit  *Audit           `json:"pendingAudit,omitempty"`
}

type RecoveryDraft struct {
	Key         string  `json:"key"`
	Content     Content `json:"content"`
	BaseVersion int     `json:"baseVersion"`
	Version     int     `json:"version"`
	UpdatedAt   string  `json:"updatedAt"`
	ExpiresAt   string  `json:"expiresAt"`
}

type draftDatabase struct {
	SchemaVersion int                      `json:"schemaVersion"`
	Sequence      int                      `json:"sequence"`
	Items         map[string]RecoveryDraft `json:"items"`
}


func loadEditorialFile(p string, limit int64, v any) error {
	// 内部业务实现略
	panic("not implemented")
}

func saveEditorialFile(p string, limit int, v any) error {
	// 内部业务实现略
	panic("not implemented")
}

func (s *Store) initEditorial() error {
	// 内部业务实现略
	panic("not implemented")
}

func cloneTaxonomy(d taxonomyDatabase) taxonomyDatabase {
	// 内部业务实现略
	panic("not implemented")
}

func (s *Store) saveTaxonomy(d taxonomyDatabase) error {
	// 内部业务实现略
	panic("not implemented")
}

func (s *Store) editorialReadyLocked() error {
	// 内部业务实现略
	panic("not implemented")
}

func taxonNameValid(name string) bool {
	// 内部业务实现略
	panic("not implemented")
}

func validateTaxonomy(d taxonomyDatabase) error {
	// 内部业务实现略
	panic("not implemented")
}

func categoryEnabled(d taxonomyDatabase, id string) bool {
	// 内部业务实现略
	panic("not implemented")
}

func (s *Store) taxonomyReferencesLocked(c *Content) error {
	// 内部业务实现略
	panic("not implemented")
}

func (s *Store) flushTaxonomyAuditLocked() error {
	// 内部业务实现略
	panic("not implemented")
}

func (s *Store) putTaxon(kind, id string, input Taxon, actor string) (Taxon, error) {
	// 内部业务实现略
	panic("not implemented")
}

func (s *Server) taxonomyRoute(w http.ResponseWriter, r *http.Request, a Actor, kind, tail string) {
	// 内部业务实现略
	panic("not implemented")
}


func validDraftKey(key string) bool {
	// 内部业务实现略
	panic("not implemented")
}

func prepareDraftContent(c Content) (Content, error) {
	// 内部业务实现略
	panic("not implemented")
}

func draftStorageKey(actor, key string) string {
	// 内部业务实现略
	panic("not implemented")
}

func activeDraft(d RecoveryDraft, now time.Time) bool {
	// 内部业务实现略
	panic("not implemented")
}

func (s *Store) putDraft(actor, key string, c Content, baseVersion, version int) (RecoveryDraft, error) {
	// 内部业务实现略
	panic("not implemented")
}

func (s *Server) draftRoute(w http.ResponseWriter, r *http.Request, a Actor, key string) {
	// 内部业务实现略
	panic("not implemented")
}
