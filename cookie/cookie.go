package cookie

import (
	"net/http"

	"go.d34d.net/pex/lg"
)

var CookieName string = "__pex_fib__"

func GetCookie(r *http.Request) {
	lg.Log.Printf("cookie.GetCookie()...")
	c, err := r.Cookie(CookieName)
	if err != nil {
		lg.Log.Printf("cookie.GetCookie() ERR:%v", err)
	}
	lg.Log.Printf("cookie.GetCookie() Cookie:%s", c)
}

func AddCookie(w http.ResponseWriter, value string) {
	cookie := http.Cookie{
		Name:  CookieName,
		Value: value,
	}
	http.SetCookie(w, &cookie)
}
