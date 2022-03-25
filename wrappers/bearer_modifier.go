package wrappers

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// SetBearerTokenFromQuery — устанавливаем Bearer токен из хеадера
func SetBearerTokenFromQuery(increq, outreq *http.Request) error {
	const (
		accessTokenKey      = "access_token"
		authorizationHeader = "Authorization"
	)

	query := increq.URL.Query()
	at, ok := query[accessTokenKey]
	if !ok {
		log.WithFields(log.Fields{"request": increq}).Warn("missing access token")
	}
	if len(at) != 0 {
		bearerToken := at[0]
		outreq.Header.Set(authorizationHeader, fmt.Sprintf("Bearer %s", bearerToken))
	}

	return nil
}
