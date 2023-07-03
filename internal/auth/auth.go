package auth

import (
	"errors"
	"net/http"
	"strings"
)

/*
GetApiKey extracts users apikey from authorization header
example auth header:
Authorization: ApiKey <apikey goes here>
*/
func GetApiKey(headers http.Header) (string, error) {
	auth_header := headers.Get("Authorization")
	if auth_header == "" {
		return "", errors.New("authorization details were not provided")
	}
	header_vals := strings.Split(auth_header, " ")
	if len(header_vals) != 2 || header_vals[0] != "ApiKey" {
		return "", errors.New("invalid authorization header")
	}
	return header_vals[1], nil
}
