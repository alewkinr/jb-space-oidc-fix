package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/alewkinr/jb-space-oidc-fix/wrappers"

	"github.com/alewkinr/jb-space-oidc-fix/config"
	log "github.com/sirupsen/logrus"
)

func main() {
	cfg := config.MustInitConfig()

	userInfoHandler := NewUserInfoHandler(cfg)

	addr := fmt.Sprintf("%s:%s", cfg.ServerHost, cfg.ServerPort)
	if runErr := http.ListenAndServe(addr, userInfoHandler); runErr != nil {
		log.Fatalf("server run err: %v", runErr)
	}
}

type UserInfoHandler struct {
	cfg             *config.Config
	issoUserInfoURL string
}

func NewUserInfoHandler(cfg *config.Config) *UserInfoHandler {
	return &UserInfoHandler{
		cfg:             cfg,
		issoUserInfoURL: cfg.IssoUserInfoURI,
	}
}

func (h *UserInfoHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	outreq, makeRequestErr := wrappers.MakeRequest(h.issoUserInfoURL, req)
	if makeRequestErr != nil {
		log.Errorf("make request: %v", makeRequestErr)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	res, doRequestErr := http.DefaultClient.Do(outreq)
	if doRequestErr != nil {
		log.Errorf("do request: %v", doRequestErr)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer func() {
		if closeBodyErr := res.Body.Close(); closeBodyErr != nil {
			log.Errorf("request close body: %v", closeBodyErr)
		}
	}()

	respBody, respReadBodyErr := ioutil.ReadAll(res.Body)
	if respReadBodyErr != nil {
		log.Errorf("read body for response: %v", respReadBodyErr)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(res.StatusCode)
	_, writeResponseBodyErr := rw.Write(respBody)
	if writeResponseBodyErr != nil {
		log.Errorf("write response: %v", writeResponseBodyErr)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
}
