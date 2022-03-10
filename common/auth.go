package common

import (
	"errors"
	"strings"

	"github.com/dgrijalva/jwt-go/request"
)

func stripBearerPrefixFromTokenString(tok string) (string, error) {
	if len(tok) == 0 {
		return "", errors.New("bearer token empty")
	} else if len(tok) > 5 && strings.ToUpper(tok[0:6]) == "TOKEN " {
		return tok[6:], nil
	} else if len(tok) <= 5 {
		return "", errors.New("invalid bearer token prefix")
	} else {
		return tok, nil
	}
}

var AuthorizationHeaderExtractor = &request.PostExtractionFilter{
	Extractor: request.HeaderExtractor{"Authorization"},
	Filter:    stripBearerPrefixFromTokenString,
}
var Auth2Extractor = &request.MultiExtractor{
	AuthorizationHeaderExtractor, request.ArgumentExtractor{"access_token"},
}
