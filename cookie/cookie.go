package cookie

import (
	"net/http"
	"strconv"

	"github.com/krsanky/pex/lg"
)

var FibIdxName string = "__pex_fib_idx__"

func GetFibIdx(r *http.Request) (idx int) {
	c, err := r.Cookie(FibIdxName)
	if err == http.ErrNoCookie {
		return // this is OK, no-cookie idx is 0
	}
	if err != nil {
		lg.Log.Printf("cookie.GetFibIdx() ERR:%v", err)
	}

	idx, err = strconv.Atoi(c.Value)
	if err != nil {
		lg.Log.Printf("cookie.GetFibIdx() Atoi:ERR:%v", err)
		idx = -1
	}

	return
}

func SetFibIdx(w http.ResponseWriter, idx int) {
	cookie := http.Cookie{
		Name:  FibIdxName,
		Value: strconv.Itoa(idx),
	}
	http.SetCookie(w, &cookie)
}
