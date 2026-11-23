// 接口定义 — source/seo.go
// 仅保留数据结构与函数签名。

package main


import (
	"bytes"
	"encoding/xml"
	"html/template"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"

	xhtml "golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)


type publicTag struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

	type entry struct {
		Loc     string `xml:"loc"`
		LastMod string `xml:"lastmod,omitempty"`
	}


func projectPublic(c Content, taxonomy taxonomyDatabase) publicContent {
	// 内部业务实现略
	panic("not implemented")
}

func visibleArticle(c Content, now time.Time) bool {
	// 内部业务实现略
	panic("not implemented")
}

func (s *Server) canonicalOrigin() (string, bool) {
	// 内部业务实现略
	panic("not implemented")
}

func renderRichBody(body string) template.HTML {
	// 内部业务实现略
	panic("not implemented")
}


func (s *Server) articleRoute(w http.ResponseWriter, r *http.Request) {
	// 内部业务实现略
	panic("not implemented")
}

func (s *Server) sitemapRoute(w http.ResponseWriter, r *http.Request) {
	// 内部业务实现略
	panic("not implemented")
}
