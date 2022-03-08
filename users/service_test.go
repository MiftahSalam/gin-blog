package users

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindOneUser(t *testing.T) {
	asserts := assert.New(t)

	for _, user := range usersMock {
		userActual, err := FindOneUser(&user)
		asserts.NoError(err, "%v should exist", user.Username)
		asserts.Equal(user, userActual, "user should equal")
	}
}
