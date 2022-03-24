package rest

import (
	"github.com/alewkinr/jb-space-oidc-fix/internal/proxy"
	"net/http"
)

// RESTAPI — структура обработчиков для REST API
type RESTAPI struct {
	reverseProxy *proxy.ReverseProxy
}

// NewRESTAPI — конструктор структуры для API
func NewRESTAPI(rp *proxy.ReverseProxy) *RESTAPI {
	return &RESTAPI{reverseProxy: rp}
}

// ProxyAuth — хэндлер для проксирования аутентификации, заменив токен из query в Header
func (r *RESTAPI) ProxyAuth(rw http.ResponseWriter, req *http.Request) {
	r.reverseProxy.ServeHTTP(rw, req)
}
