package pex

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/pelletier/go-toml"
	"go.d34d.net/pex/api"
	"go.d34d.net/pex/cookie"
	"go.d34d.net/pex/db"
	"go.d34d.net/pex/lg"
	"go.d34d.net/pex/server"
	"go.d34d.net/pex/session"
)

func testIndex(w http.ResponseWriter, r *http.Request) {
	cookie.AddCookie(w, "asd asd")
	w.Header().Set("Content-Type", "application/json")
	cookie.GetCookie(r)

	// test session
	skey := "pex-api-test-session-key"
	sstr := session.Session.GetString(r.Context(), skey)
	lg.Log.Printf("sstr:%s", sstr)
	if sstr == "" {
		session.Session.Put(r.Context(), skey, "1")
	} else {
		i, err := strconv.Atoi(sstr)
		if err != nil {
			lg.Log.Printf("pex.Index ERR:%v", err)
		} else {
			i += 1
			session.Session.Put(r.Context(), skey, fmt.Sprintf("%d", i))
		}
	}


	mp := make(map[string]interface{})
	mp["pex-fib"] = sstr
	bs, _ := json.Marshal(mp)

	w.Write(bs)
}

func Web(settings *toml.Tree) {
	db.Init(settings)
	mux := http.NewServeMux()
	mux.HandleFunc("/", testIndex)
	api.AddRoutes(mux)
	server := server.NewServer(settings, mux)

	server.AddSessionHandler(db.DB)

	fmt.Println("server.ServeDev()...")
	server.ServeDev()
}
