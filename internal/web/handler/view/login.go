package view

import (
	"net/http"

	"github.com/hidu/airshare/internal/sharesrv"
	"github.com/hidu/airshare/internal/web/tpls"
)

func login(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	token := query.Get("token")
	values := map[string]any{}
	if token != "" {
		if token == sharesrv.Token() {
			http.Redirect(w, r, "j?token="+token, http.StatusFound)
			return
		}
		values["Error"] = "Invalid Token"
	}
	tpls.RenderFrontPage(w, r, "view/login.html", values)
}
