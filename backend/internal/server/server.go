package server

import (
	"net/http"
)

type Server struct {
	Server *http.Server
}

func (s *Server) Run(port string, handler http.Handler) error {
	s.Server = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
	}
	//log.Printf("Server running on port %s...", s.Server.Addr)
	return s.Server.ListenAndServe()
}
