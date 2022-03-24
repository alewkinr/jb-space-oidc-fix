package proxy

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"net/http/httputil"
	"net/url"
)

// ReverseProxy — реверс прокси
type ReverseProxy struct {
	host      string
	httpProxy *httputil.ReverseProxy
}

// NewReverseProxy — реверс прокси для аутентификации по OIDC
func NewReverseProxy(targetHost string) (*ReverseProxy, error) {
	proxy := &ReverseProxy{host: targetHost}
	if setUpErr := proxy.setUp(); setUpErr != nil {
		return nil, setUpErr
	}

	return proxy, nil
}

func (p *ReverseProxy) setUp() error {
	u, urlParseErr := url.Parse(p.host)
	if urlParseErr != nil {
		return urlParseErr
	}

	proxy := httputil.NewSingleHostReverseProxy(u)
	originalDirector := proxy.Director
	proxy.Director = func(req *http.Request) {
		originalDirector(req)

		// устанавливаем токен из query в Authorization Header
		setBearerTokenFromQuery(req)
	}

	proxy.ErrorHandler = p.errorHandler
	p.httpProxy = proxy
	return nil
}

// errorHandler — обработчик ошибок проксирования
func (p *ReverseProxy) errorHandler(rw http.ResponseWriter, req *http.Request, err error) {
	if err != nil {
		log.WithFields(log.Fields{"request": req, "response": rw, "error": err}).Error("reverse proxy err")
	}

	rw.WriteHeader(http.StatusBadGateway)
	return
}

// ServeHTTP — реализуем метод ServHTTP
func (p *ReverseProxy) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if p.httpProxy == nil {
		log.Warn("proxy must init before use")
		rw.WriteHeader(http.StatusBadGateway)
		return
	}

	p.httpProxy.ServeHTTP(rw, r)
}
