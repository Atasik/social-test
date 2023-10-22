package server

import (
	"context"
	"net/http"
	"social/internal/config"
)

type server struct {
	httpServer *http.Server
}

func NewServer(cfg *config.Config, handler http.Handler) *server {
	return &server{
		httpServer: &http.Server{
			Addr:           ":" + cfg.HTTP.Port,
			Handler:        handler,
			MaxHeaderBytes: cfg.HTTP.MaxHeaderMegaBytes << 18, //nolint:gomnd
			ReadTimeout:    cfg.HTTP.ReadTimeout,
			WriteTimeout:   cfg.HTTP.WriteTimeout,
		},
	}
}

func (s *server) Run() error {
	return s.httpServer.ListenAndServe()
}
func (s *server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
