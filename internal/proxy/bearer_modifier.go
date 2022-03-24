package proxy

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
)

// setBearerTokenFromQuery — устанавливаем Bearer токен из хеадера
func setBearerTokenFromQuery(req *http.Request) {
	const (
		accessTokenKey      = "access_token"
		authorizationHeader = "Authorization"
	)

	query := req.URL.Query()
	at, ok := query[accessTokenKey]
	if !ok {
		log.WithFields(log.Fields{"request": req}).Warn("missing access token")
	}
	if len(at) != 0 {
		bearerToken := at[0]
		req.Header.Set(authorizationHeader, fmt.Sprintf("Bearer %s", bearerToken))
	}
}
