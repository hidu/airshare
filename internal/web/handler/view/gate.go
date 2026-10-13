package view

import (
	"net/http"

	"github.com/xanygo/anygo/xhttp"

	"github.com/hidu/airshare/internal/sharesrv"
)

const cookieName = "token"

func jump(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	token := query.Get("token")
	if token != sharesrv.Token() {
		xhttp.Error(w, r, http.StatusForbidden, "Forbidden", "Forbidden")
		return
	}
	ck := &http.Cookie{
		Name:     cookieName,
		Value:    token,
		Path:     sharesrv.SharePath(),
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(w, ck)
	http.Redirect(w, r, sharesrv.SharePath(), http.StatusFound)
}

func auth() xhttp.MiddlewareFunc {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			cookie, _ := r.Cookie(cookieName)
			if cookie == nil || cookie.Value != sharesrv.Token() {
				xhttp.Error(w, r, http.StatusForbidden, "Forbidden", "Forbidden")
				return
			}
			h.ServeHTTP(w, r)
		})
	}
}
