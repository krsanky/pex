package server

import (
	"fmt"
	"net"
	"net/http"
	"net/http/fcgi"

	"github.com/pelletier/go-toml"
)

type Server struct {
	Settings *toml.Tree
	Handler  http.Handler
}

func NewServer(settings *toml.Tree, h http.Handler) *Server {
	s := Server{}
	s.Settings = settings
	s.Handler = h
	return &s
}

// Add routes to mux before calling this
/*
func (s *Server) SetupHandler() {
	// ORDER MATTERS and it's kind of reversed
	h := nosurf.NewPure(s.Handler)
	h = account.AddUser(h)
	//session.Init()
	//h = session.Session.LoadAndSave(h)
	s.Handler = h
}
*/

func (s *Server) Serve() {
	port := s.Settings.Get("server.port").(int64)
	ip := s.Settings.Get("server.ip").(string)

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", ip, port))
	if err != nil {
		panic(err)
	}
	fcgi.Serve(listener, s.Handler)
}

/*
rv := &http.Server{
    ReadTimeout: 5 * time.Second,
    WriteTimeout: 10 * time.Second,
}
log.Println(srv.ListenAndServe())
*/

func (s *Server) ServeDev() {
	port := s.Settings.Get("server.port").(int64)
	ip := s.Settings.Get("server.ip").(string)

	http_server := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", ip, port),
		Handler: s.Handler,
	}
	http_server.ListenAndServe()
}
