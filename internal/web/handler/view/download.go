package view

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/xanygo/anygo/xhttp"
)

func downloadForDir(w http.ResponseWriter, r *http.Request, dir string) {
	query := r.URL.Query()
	name := query.Get("name")
	if name == "" {
		xhttp.Error(w, r, 400, "文件不存在", "文件不存在")
		return
	}
	fp := filepath.Join(dir, name)

	// 防止访问分享目录之外的文件
	rel, err := filepath.Rel(dir, fp)
	if err != nil || rel == ".." || strings.HasPrefix(rel, ".."+string(filepath.Separator)) {
		xhttp.Error(w, r, 403, "禁止访问", "禁止访问")
		return
	}

	info, err := os.Stat(fp)
	if err != nil || !info.Mode().IsRegular() {
		xhttp.Error(w, r, 404, "文件不存在", "文件不存在")
		return
	}

	if query.Get("action") == "download" {
		cd := fmt.Sprintf("attachment; filename*=UTF-8''%s", url.PathEscape(filepath.Base(name)))
		w.Header().Set("Content-Disposition", cd)
	}

	http.ServeFile(w, r, fp)
}
