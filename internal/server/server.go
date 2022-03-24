package server

import (
	"context"
	"fmt"
	"github.com/alewkinr/jb-space-oidc-fix/internal/config"
	"github.com/alewkinr/jb-space-oidc-fix/internal/server/rest"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type HTTPServer struct {
	s *http.Server
}

// NewHTTPServer — http сервер
func NewHTTPServer(cfg *config.Config, rest *rest.RESTAPI) *HTTPServer {
	router := http.NewServeMux()
	router.HandleFunc("/", rest.ProxyAuth)

	return &HTTPServer{
		s: &http.Server{
			Addr:    fmt.Sprintf("%s:%s", cfg.ServerHost, cfg.ServerPort),
			Handler: router,
		},
	}
}

// Run запускает сервер
func (ps *HTTPServer) Run() error {
	return ps.s.ListenAndServe()
}

// Stop останавливает сервер
func (ps *HTTPServer) Stop(ctx context.Context) error {
	log.Info("Stopping public server")
	return ps.s.Shutdown(ctx)
}
