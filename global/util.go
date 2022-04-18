package global

import (
	"github.com/gin-gonic/gin"
	"strings"
)

type Utils struct {
}

func (Utils) Eq(a, b string) bool {
	if strings.Compare(a, b) == 0 {
		return true
	}
	return false
}
func (Utils) CurrentUserId(ctx *gin.Context) {
	value, _ := ctx.Get(CONF.Jwt.Context)
	s := value.(int)
	println("s", s)
}
