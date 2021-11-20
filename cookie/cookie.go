package cookie

import (
	"net/http"
	"strconv"

	"go.d34d.net/pex/lg"
)

var FibIdxName string = "__pex_fib_idx__"

func GetFibIdx(r *http.Request) (idx int) {
	c, err := r.Cookie(FibIdxName)
	if err == http.ErrNoCookie {
		return // default idx is 0
	}
	if err != nil {
		lg.Log.Printf("cookie.GetCookie() ERR:%v", err)
	}
	lg.Log.Printf("cookie.GetCookie() Cookie:%s", c)

	idx, err = strconv.Atoi(c.Value)
	if err != nil {
		lg.Log.Printf("cookie.GetCookie() Atoi:ERR:%v", err)
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
