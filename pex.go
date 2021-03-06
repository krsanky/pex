package pex

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/pelletier/go-toml"
	"github.com/krsanky/pex/api"
	"github.com/krsanky/pex/server"
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

	fmt.Printf("server.Serve() env:%s port:%d\n", settings.Get("server.env").(string), settings.Get("server.port").(int64))
	server.Serve()
}
