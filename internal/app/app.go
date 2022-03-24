package app

import (
	"github.com/alewkinr/jb-space-oidc-fix/internal/config"
	proxy "github.com/alewkinr/jb-space-oidc-fix/internal/proxy"
	"github.com/alewkinr/jb-space-oidc-fix/internal/server"
	"github.com/alewkinr/jb-space-oidc-fix/internal/server/rest"
	log "github.com/sirupsen/logrus"
)

// App — простецкое приложение
type App struct {
	server *server.HTTPServer
}

// MustNewApp — инициализируем приложение или паникуем
func MustNewApp() *App {
	app := &App{}
	if setUpErr := app.setUp(); setUpErr != nil {
		panic(setUpErr)
	}

	return app
}

func (a *App) setUp() error {
	cfg := config.MustInitConfig()

	reverseProxy, initReverseProxyErr := proxy.NewReverseProxy(cfg.ProviderHost)
	if initReverseProxyErr != nil {
		log.Errorf("init proxy: %v", initReverseProxyErr)
		return initReverseProxyErr
	}

	api := rest.NewRESTAPI(reverseProxy)
	a.server = server.NewHTTPServer(cfg, api)
	return nil
}

// Run — запускаем приложение
func (a *App) Run() error {
	return a.server.Run()
}
