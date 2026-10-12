package admin

import (
	"github.com/xanygo/anygo/xhttp"
)

func Router(r *xhttp.Router) {
	r.GetFunc("/", index)
}
