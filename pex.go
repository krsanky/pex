package pex

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/pelletier/go-toml"
	"go.d34d.net/pex/api"
	"go.d34d.net/pex/server"
)

func index(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl.Execute(w, nil)
}

func Web(settings *toml.Tree) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	api.AddRoutes(mux)
	server := server.NewServer(settings, mux)

	fmt.Printf("server.ServeDev() port:%d\n", settings.Get("server.port").(int64))
	server.ServeDev()
}
