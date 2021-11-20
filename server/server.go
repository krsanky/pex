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

func (s *Server) Serve() {
	port := s.Settings.Get("server.port").(int64)
	ip := s.Settings.Get("server.ip").(string)

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", ip, port))
	if err != nil {
		panic(err)
	}

	/*
	   srv := &http.Server{
	       ReadTimeout: 5 * time.Second,
	       WriteTimeout: 10 * time.Second,
	   }
	   srv.ListenAndServe()
	*/

	fcgi.Serve(listener, s.Handler)
}

func (s *Server) ServeDev() {
	port := s.Settings.Get("server.port").(int64)
	ip := s.Settings.Get("server.ip").(string)

	http_server := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", ip, port),
		Handler: s.Handler,
	}
	http_server.ListenAndServe()
}
