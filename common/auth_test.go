package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStripBearerPrefixFromTokenString(t *testing.T) {
	asserts := assert.New(t)
	var bearerString = "ashdoawuriwjsskfn"
	tableTest := []struct {
		name string
		args string
		want string
	}{
		{"No error test: string without prefix", bearerString, bearerString},
		{"No error test: string with capitalized prefix", "TOKEN " + bearerString, bearerString},
		{"No error test: string with uncapitalized prefix", "token " + bearerString, bearerString},
	}

	for _, test := range tableTest {
		t.Run(test.name, func(t *testing.T) {
			tok, err := stripBearerPrefixFromTokenString(test.args)
			asserts.NoError(err, "should return non empty string")
			asserts.Equal(test.want, tok, "returned string should be same")
		})
	}

	tableTest = []struct {
		name string
		args string
		want string
	}{
		{"Error test: string empty", "", ""},
		{"Error test: string invalid", "sdf", ""},
	}
	for _, test := range tableTest {
		t.Run(test.name, func(t *testing.T) {
			tok, err := stripBearerPrefixFromTokenString(test.args)
			asserts.Error(err, "should return empty string")
			asserts.Equal(test.want, tok, "returned string should be empty")
		})
	}
}
