package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/pelletier/go-toml"
	"go.d34d.net/pex/server"
)

func usage() {
	fmt.Printf("pex <settings.toml>\n")
}

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	mp := make(map[string]interface{})
	mp["pex"] = true
	bs, _ := json.Marshal(mp)
	w.Write(bs)
}

func main() {
	if len(os.Args) == 2 {
		sfile := os.Args[1]
		settings, err := toml.LoadFile(sfile)
		if err != nil {
			panic(err)
		}

		//db.Init(settings)

		mux := http.NewServeMux()
		mux.HandleFunc("/", Index)
		server := server.NewServer(settings, mux)

		server.ServeDev()
	} else {
		usage()
	}
}
