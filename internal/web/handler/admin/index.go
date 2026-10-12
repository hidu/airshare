package admin

import (
	"net/http"

	"github.com/hidu/airshare/internal/sharesrv"
	"github.com/hidu/airshare/internal/web/tpls"
)

func index(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	reflush := query.Get("reflush")
	switch reflush {
	case "long":
		sharesrv.FlushLongToken()
		http.Redirect(w, r, r.URL.Path, http.StatusFound)
		return
	case "short":
		sharesrv.FlushShortToken()
		http.Redirect(w, r, r.URL.Path, http.StatusFound)
		return
	default:
	}
	base := sharesrv.ShareSite() + sharesrv.SharePath()
	values := map[string]any{
		"Token":     sharesrv.Token(),
		"ShareDir":  sharesrv.ShareDir(),
		"ShareSite": sharesrv.ShareSite(),
		"URL":       base + "j?token=" + sharesrv.Token(),
	}
	tpls.RenderAdminPage(w, r, "admin/index.html", values)
}
