package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"go.d34d.net/pex/lg"
	"go.d34d.net/pex/session"
)

func AddRoutes(m *http.ServeMux) {
	m.HandleFunc("/api", testIndex)
}

func testIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// test session
	skey := "pex-fib"
	sstr := session.Session.GetString(r.Context(), skey)
	lg.Log.Printf("sstr:%s", sstr)
	if sstr == "" {
		session.Session.Put(r.Context(), skey, "1")
	} else {
		i, err := strconv.Atoi(sstr)
		if err != nil {
			lg.Log.Printf("pex.Index ERR:%v", err)
		} else {
			i += i
			session.Session.Put(r.Context(), skey, fmt.Sprintf("%d", i))
		}
	}

	mp := make(map[string]interface{})
	mp["pex:testIndex"] = true
	bs, _ := json.Marshal(mp)

	w.Write(bs)
}
