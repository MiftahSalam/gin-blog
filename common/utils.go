package common

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func RandString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func GetToken(id uint) string {
	jwt_token := jwt.New(jwt.GetSigningMethod("HS256"))
	jwt_expired_env := os.Getenv("JWT_EXPIRED_IN")
	jwt_expired, err := strconv.Atoi(jwt_expired_env)

	if err != nil {
		log.Fatal("Error while loading jwt_expired. Err: ", err)
	}

	jwt_token.Claims = jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(time.Minute * time.Duration(jwt_expired)),
	}
	token, _ := jwt_token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	return token
}

type CommonError struct {
	Errors map[string]interface{} `json:"errors"`
}

func NewValidatorError(err error) CommonError {
	res := CommonError{}
	res.Errors = make(map[string]interface{})
	errs := err.(validator.ValidationErrors)

	for _, v := range errs {
		if v.Param() != "" {
			res.Errors[v.Field()] = fmt.Sprintf("{%v: %v}", v.Tag(), v.Param())
		} else {
			res.Errors[v.Field()] = fmt.Sprintf("{key: %v}", v.Tag())
		}
	}

	return res
}
