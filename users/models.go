package users

import (
	"fmt"

	"github.com/MiftahSalam/gin-blog/common"
)

func CheckDotEnv() {
	fmt.Println(common.GetToken(23))
}
