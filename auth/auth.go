package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetAPIKey(headers http.Header) (string, error) {
	auth := headers.Get("Authorization")
	if auth == "" {
		return "", errors.New("no authorization headers where provided!")
	}

	value := strings.Split(auth, " ")
	if len(value) < 2 || value[0] != "ApiKey" {
		return "", errors.New("invalid authorization headers")
	}

	return value[1], nil
}

