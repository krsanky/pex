package pex

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pelletier/go-toml"
	"go.d34d.net/pex/api"
	"go.d34d.net/pex/server"
)

func testIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	mp := make(map[string]interface{})
	mp["pex-fib"] = "12312312"
	bs, _ := json.Marshal(mp)

	w.Write(bs)
}

func Web(settings *toml.Tree) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", testIndex)
	api.AddRoutes(mux)
	server := server.NewServer(settings, mux)

	fmt.Println("server.ServeDev()...")
	server.ServeDev()
}
