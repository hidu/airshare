package main

import (
	"flag"

	"github.com/xanygo/anygo/xattr"
	"github.com/xanygo/anygo/xcfg"
	"github.com/xanygo/ext"

	"github.com/hidu/airshare/internal/bootstrap"
	"github.com/hidu/airshare/internal/web"
)

func init() {
	ext.Init()
}

var c = flag.String("conf", "conf/app.yml", "app main config file")

func main() {
	flag.Parse()
	xattr.MustInitAppMain(*c, xcfg.Parse)
	bootstrap.Bootstrap()
	web.Run()
}
