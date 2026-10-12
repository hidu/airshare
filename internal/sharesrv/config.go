package sharesrv

import (
	"fmt"
	"path"
	"path/filepath"
	"strings"

	"github.com/xanygo/anygo/xattr"
	"github.com/xanygo/anygo/xsync"
)

func DownloadDir() string {
	return filepath.Join(ShareDir(), "download")
}
func UploadDir() string {
	return filepath.Join(ShareDir(), "upload")
}

func ShareDir() string {
	value, err := xattr.GetOtherAs[string]("ShareDir")
	if err != nil {
		panic(fmt.Errorf("read  ShareDir failed: %w", err))
	}
	return value
}

var sharePath = &xsync.OnceInit[string]{
	New: func() string {
		sharePath := "/" + xattr.GetOtherDefault("SharePath", "/view/")
		sharePath = path.Clean(sharePath) + "/"
		return sharePath
	},
}

func SharePath() string {
	return sharePath.Load()
}

func ShareSite() string {
	site := xattr.GetOtherDefault("ShareSite", "")
	if site == "" {
		panic("empty ShareSite")
	}
	return strings.TrimRight(site, "/")
}
