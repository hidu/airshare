package web

import (
	"context"
	"net/http"

	"github.com/xanygo/anygo/xattr"
	"github.com/xanygo/anygo/xlog"
	"github.com/xanygo/anygo/xmime"
)

func Run() {
	routerRegister()
	xmime.Register()

	listen := xattr.AppMain().MustGetListen("main")
	xlog.Debug(context.Background(), "server Listen", xlog.String("Listen", listen))
	ser := &http.Server{
		Handler: router,
		Addr:    listen,
	}
	ser.SetKeepAlivesEnabled(true)
	err := ser.ListenAndServe()
	xlog.Warn(context.Background(), "server exit", xlog.Err("error", err))
}
