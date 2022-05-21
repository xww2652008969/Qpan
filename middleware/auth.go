package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Jwtauth() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("不让你通过，略略略")
		c.Abort()
	}
}
