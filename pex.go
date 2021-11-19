package pex

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pelletier/go-toml"
	"go.d34d.net/pex/db"
	"go.d34d.net/pex/server"
)

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
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

	fmt.Println("server.ServeDev()...")
	server.ServeDev()
}
