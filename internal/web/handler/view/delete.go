package view

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/hidu/airshare/internal/sharesrv"
	"github.com/xanygo/webr"
)

func uDelete(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	name := query.Get("name")
	if name == "" {
		webr.WriteJSON(w, 400, "invalid request", nil)
		return
	}
	dir := sharesrv.UploadDir()
	fp := filepath.Join(dir, name)

	rel, err := filepath.Rel(dir, fp)
	if err != nil || rel == ".." || strings.HasPrefix(rel, ".."+string(filepath.Separator)) {
		webr.WriteJSON(w, 403, "invalid name", nil)
		return
	}

	info, err := os.Stat(fp)
	if err != nil || !info.Mode().IsRegular() {
		webr.WriteJSON(w, 404, "file not found", nil)
		return
	}
	err = os.Remove(fp)
	webr.WriteJSONAuto(w, err)
}
