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
)

var te *xhtml.TemplateEngine

func init() {
	maps.Copy(funcmap, xhtml.FuncMap)
	t := template.Must(xhtml.WalkParseFS(template.New("files").Funcs(funcmap), files, "files", "*.html"))
	te = &xhtml.TemplateEngine{
		Temp: t,
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
