package web

import (
	"net/http"

	"github.com/xanygo/anygo/xhttp"
	"github.com/xanygo/anygo/xhttp/xhandler"
	"github.com/xanygo/anygo/xi18n"
	"github.com/xanygo/anygo/xlog"

	"github.com/hidu/airshare/internal/bootstrap"
	"github.com/hidu/airshare/internal/web/handler"
)

var router = xhttp.NewRouter()

func routerRegister() {
	al := &xhandler.AccessLog{
		Logger: xlog.AccessLogger(),
	}
	router.Use(al.Next)

	xi := &xi18n.HTTPHandler{
		Bundle: bootstrap.Bundle,
	}
	router.Use(xi.Next)

	// router.GetFunc("/favicon.ico", favicon)
	router.GetFunc("/", index)
	handler.Router(router)
}

func index(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/admin/", http.StatusFound)
}
