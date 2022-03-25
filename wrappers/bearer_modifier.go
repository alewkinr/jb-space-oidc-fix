package wrappers

import (
	"fmt"
	"net/http"
	"net/url"

	log "github.com/sirupsen/logrus"
)

// SetBearerTokenFromQuery — устанавливаем Bearer токен из хеадера
func SetBearerTokenFromQuery(sourceURL string, req *http.Request) error {
	const (
		accessTokenKey      = "access_token"
		authorizationHeader = "Authorization"
	)

	sourceRequestURL, parseURLErr := url.Parse(sourceURL)
	if parseURLErr != nil {
		return parseURLErr
	}

	query := sourceRequestURL.Query()
	at, ok := query[accessTokenKey]
	if !ok {
		log.WithFields(log.Fields{"request": req}).Warn("missing access token")
	}
	if len(at) != 0 {
		bearerToken := at[0]
		req.Header.Set(authorizationHeader, fmt.Sprintf("Bearer %s", bearerToken))
	}

	return nil
}
