package pex

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/pelletier/go-toml"
	"go.d34d.net/pex/db"
	"go.d34d.net/pex/lg"
	"go.d34d.net/pex/server"
	"go.d34d.net/pex/session"
)

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

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
	mp["pex"] = true
	bs, _ := json.Marshal(mp)

	w.Write(bs)
}

func API(settings *toml.Tree) {
	db.Init(settings)

	mux := http.NewServeMux()
	mux.HandleFunc("/", Index)
	server := server.NewServer(settings, mux)
	server.AddSessionHandler(db.DB)

	fmt.Println("server.ServeDev()...")
	server.ServeDev()
}
