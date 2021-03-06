package common

import (
	"log"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.M) {
	LogI.Println("Test main users common/utils start")

	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Cannot load env file. Err: ", err)
	}

	t.Run()

	LogI.Println("Test main users common/utils end")
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
	token := GetToken(3)

	asserts.IsType(token, string("token"), "token type should be string")
}
