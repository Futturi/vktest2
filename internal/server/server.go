package server

import (
	"context"
	"net/http"
	"time"
)

type Server struct {
	Server *http.Server
}

func (s *Server) InitServer(port string, handler http.Handler) error {
	s.Server = &http.Server{
		Handler:      handler,
		Addr:         ":" + port,
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}
	return s.Server.ListenAndServe()
}
func (s *Server) StopServer(ctx context.Context) error {
	return s.Server.Shutdown(ctx)
}
