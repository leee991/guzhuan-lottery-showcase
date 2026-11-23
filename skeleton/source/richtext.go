// 接口定义 — source/richtext.go
// 仅保留数据结构与函数签名。


package main


import (
	"bytes"
	"errors"
	"io"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/microcosm-cc/bluemonday"
	xhtml "golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)


func validateEditorialFields(c Content) error {
	// 内部业务实现略
	panic("not implemented")
}

func prepareContent(c Content) (Content, error) {
	// 内部业务实现略
	panic("not implemented")
}

func sanitizeHTML(raw string) (string, error) {
	// 内部业务实现略
	panic("not implemented")
}

func bodyText(c Content) string {
	// 内部业务实现略
	panic("not implemented")
}
