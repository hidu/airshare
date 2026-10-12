package web

import (
	"github.com/hidu/airshare/internal/web/handler"
	"github.com/xanygo/anygo/xhttp"
)

var router = xhttp.NewRouter()

func routerRegister() {
	// router.GetFunc("/favicon.ico", favicon)
	handler.Router(router)
}
