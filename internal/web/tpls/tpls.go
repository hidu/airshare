//  Copyright(C) 2025 github.com/hidu  All Rights Reserved.
//  Author: hidu <duv123+git@gmail.com>
//  Date: 2025-09-19

package tpls

import (
	_ "embed"
	"html/template"
	"maps"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/xanygo/anygo/xhtml"
	"github.com/xanygo/anygo/xi18n"

	"github.com/hidu/airshare/internal/bootstrap"
)

var te *xhtml.TemplateEngine

func init() {
	maps.Copy(funcmap, xhtml.FuncMap)
	it := xi18n.BuildTemplate(xi18n.LangZh, bootstrap.Bundle, func(t *template.Template) *template.Template {
		return template.Must(xhtml.WalkParseFS(t.Funcs(funcmap), files, "files", "*.html"))
	})

	te = &xhtml.TemplateEngine{
		TemplateResolver: func(r *http.Request) xhtml.TemplateExecutor {
			languages := xi18n.LanguagesFromContext(r.Context())
			return it.Any(xi18n.LangZh, languages...)
		},
	}
}

func Render(w http.ResponseWriter, r *http.Request, fileName string, values map[string]any) {
	te.Render(w, r, fileName, values)
}

func RenderAdminPage(w http.ResponseWriter, r *http.Request, fileName string, values map[string]any) {
	te.RenderPageLayout(w, r, fileName, "layout/admin.html", values)
}

func RenderFrontPage(w http.ResponseWriter, r *http.Request, fileName string, values map[string]any) {
	te.RenderPageLayout(w, r, fileName, "layout/front.html", values)
}

var funcmap = template.FuncMap{
	"isImage": isImage,
}

// IsImage 判断文件名是否为图片类型。
func isImage(filename string) bool {
	switch strings.ToLower(filepath.Ext(filename)) {
	case ".jpg", ".jpeg", ".png", ".gif", ".webp",
		".bmp", ".svg", ".ico", ".tif", ".tiff",
		".avif", ".heic", ".heif":
		return true
	default:
		return false
	}
}
