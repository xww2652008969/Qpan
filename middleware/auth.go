package middleware

import (
	"Qpan/model"
	"Qpan/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Jwtauth() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := c.GetHeader("token")
		name, err := utils.Parsetoken(t)
		if err == nil {
			c.Request.Header.Add("name", name)
			c.Next()
		} else {
			fmt.Println("失效")
			c.Abort()
			model.FailWithMessage(c.Request.Header, "验证失败", c)
		}
	}
}
