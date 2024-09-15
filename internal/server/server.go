package server

import (
	"context"
	"fmt"
	"net/http"
	"schoolGradeCalculator/internal/config"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func New(cfg *config.Config, handler http.Handler) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:           fmt.Sprintf(":%d", cfg.Http.Port),
			Handler:        handler,
			MaxHeaderBytes: 1 << 20, // 1 MB
			ReadTimeout:    60 * time.Second,
			WriteTimeout:   300 * time.Second,
		},
	}
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
