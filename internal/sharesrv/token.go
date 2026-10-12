package sharesrv

import (
	"math/rand/v2"
	"strconv"
	"time"

	"github.com/xanygo/anygo/xstr"
	"github.com/xanygo/anygo/xsync"
)

var token = &xsync.OnceInit[string]{
	New: createToken,
}

func Token() string {
	return token.Load()
}

func createToken() string {
	nt := time.Now().Format("01021504") + "-" + xstr.RandNChar(6)
	return nt
}

func FlushLongToken() string {
	nt := createToken()
	token.Store(nt)
	return nt
}
func FlushShortToken() string {
	nt := strconv.Itoa(rand.Int() % 10000)
	token.Store(nt)
	return nt
}
