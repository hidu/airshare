package view

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/hidu/airshare/internal/sharesrv"
	"github.com/xanygo/anygo/xctx"
	"github.com/xanygo/anygo/xio/xfs"
	"github.com/xanygo/webr"
)

func upload(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(30 << 20)
	if err != nil {
		webr.WriteJSONError(w, err)
		return
	}
	if r.MultipartForm != nil {
		defer r.MultipartForm.RemoveAll()
	}

	mf, mr, err := r.FormFile("file")
	err = xctx.CheckError(r.Context(), err)
	if err != nil {
		webr.WriteJSONError(w, err)
		return
	}
	defer mf.Close()

	// 获取文件名，防止路径穿越
	name := filepath.Base(mr.Filename)

	if name == "." || name == "" {
		webr.WriteJSONErrorStatus(w, fmt.Errorf("文件名无效"), http.StatusBadRequest)
		return
	}

	dir := sharesrv.UploadDir()
	_ = os.Mkdir(dir, 0777)
	fp := filepath.Join(dir, name)
	for i := 1; ; i++ {
		exists, err := xfs.Exists(fp)
		if err != nil {
			webr.WriteJSONErrorStatus(w, fmt.Errorf("系统错误：%w", err), http.StatusInternalServerError)
			return
		}
		if !exists {
			break
		}
		fp = filepath.Join(dir, fmt.Sprintf("%d_%s", i, name))
	}

	out, err := os.OpenFile(fp, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0644)
	if err != nil {
		webr.WriteJSONErrorStatus(w, fmt.Errorf("上传失败，无法创建文件：%w", err), http.StatusInternalServerError)
		return
	}
	defer out.Close()

	if _, err = io.Copy(out, mf); err != nil {
		_ = os.Remove(fp)
		webr.WriteJSONErrorStatus(w, fmt.Errorf("保存文件失败：%w", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	webr.WriteJSON(w, 200, "success", nil)
}
