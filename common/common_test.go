package common

import (
	"log"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.M) {
	err := godotenv.Load("../.env")

	if err != nil {
		log.Fatal("Cannot load env file. Err: ", err)
	}
}
func TestRandString(t *testing.T) {
	asserts := assert.New(t)

	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	str := RandString(0)
	asserts.Equal(str, "", "Length should be 0")

	str = RandString(10)
	asserts.Equal(len(str), 10, "length should be 10")

	for _, ch := range str {
		asserts.Contains(letters, ch, "char should contain")
	}
}

func TestGenToken(t *testing.T) {
	asserts := assert.New(t)
	token := GetToken(2)

	asserts.IsType(token, string("token"), "token type should be string")
	asserts.Len(token, 115, "Jwt's length should be 115")
}
