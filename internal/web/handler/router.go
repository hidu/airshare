package handler

import (
	"github.com/hidu/airshare/internal/sharesrv"
	"github.com/hidu/airshare/internal/web/handler/admin"
	"github.com/hidu/airshare/internal/web/handler/view"
	"github.com/xanygo/anygo/xhttp"
)

func Router(r *xhttp.Router) {
	admin.Router(r.Prefix("/admin/"))
	v := r.Prefix(sharesrv.SharePath())
	view.Router(v)
}
