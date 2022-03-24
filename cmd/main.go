package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/alewkinr/jb-space-oidc-fix/config"
	"github.com/alewkinr/jb-space-oidc-fix/wrappers"
	log "github.com/sirupsen/logrus"
)

func main() {
	cfg := config.MustInitConfig()

	redirectURL, parseURLErr := url.Parse(cfg.ProviderHost)
	if parseURLErr != nil {
		log.Fatalf("parse redirect URL: %v", parseURLErr)
	}

	proxy := httputil.NewSingleHostReverseProxy(redirectURL)
	originalDirector := proxy.Director
	proxy.Director = func(req *http.Request) {
		originalDirector(req)

		// устанавливаем токен из query в Authorization Header
		wrappers.SetBearerTokenFromQuery(req)
	}

	addr := fmt.Sprintf("%s:%s", cfg.ServerHost, cfg.ServerPort)
	if runErr := http.ListenAndServe(addr, proxy); runErr != nil {
		log.Fatalf("server run err: %v", runErr)
	}
}
