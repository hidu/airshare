package view

import (
	"net/http"

	"github.com/hidu/airshare/internal/sharesrv"
	"github.com/xanygo/anygo/xhttp"
)

func Router(r *xhttp.Router) {
	r.GetFunc("/j", jump)
	r.GetFunc("/l", login)
	r.Use(auth())
	r.GetFunc("/", index)
	r.GetFunc("/d-file", func(w http.ResponseWriter, r *http.Request) {
		downloadForDir(w, r, sharesrv.DownloadDir())
	})
	r.GetFunc("/u-file", func(w http.ResponseWriter, r *http.Request) {
		downloadForDir(w, r, sharesrv.UploadDir())
	})
	r.PostFunc("/upload", upload)
	r.DeleteFunc("/u-delete", uDelete)
}
