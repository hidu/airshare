package view

import (
	"errors"
	"io/fs"
	"net/http"
	"path/filepath"
	"slices"

	"github.com/hidu/airshare/internal/sharesrv"
	"github.com/hidu/airshare/internal/web/tpls"
	"github.com/xanygo/anygo/xcmp"
	"github.com/xanygo/anygo/xhttp"
)

func index(w http.ResponseWriter, r *http.Request) {
	downloadFiles, err := listFiles(sharesrv.DownloadDir())
	if err != nil {
		xhttp.Error(w, r, http.StatusInternalServerError, "Error", err.Error())
		return
	}

	uploadFiles, err := listFiles(sharesrv.UploadDir())
	if err != nil {
		xhttp.Error(w, r, http.StatusInternalServerError, "Error", err.Error())
		return
	}

	values := map[string]any{
		"DownloadFiles": downloadFiles,
		"UploadFiles":   uploadFiles,
	}
	tpls.RenderFrontPage(w, r, "view/index.html", values)
}

func listFiles(dir string) ([]fs.DirEntry, error) {
	var files []fs.DirEntry
	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err == nil && !d.IsDir() {
			files = append(files, d)
		}
		return err
	})
	slices.SortStableFunc(files, xcmp.OrderAsc(func(d fs.DirEntry) int64 {
		info, err := d.Info()
		if err != nil {
			return 0
		}
		return info.ModTime().UnixNano()
	}))
	if err != nil && errors.Is(err, fs.ErrNotExist) {
		err = nil
	}
	return files, err
}
