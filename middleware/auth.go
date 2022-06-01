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
		uuid, err := utils.Parsetoken(t)
		if err == nil {
			c.Request.Header.Add("uuid", fmt.Sprintf("%v", uuid))
			out, err := utils.Getredis(fmt.Sprintf("%v", uuid), "jwt")
			fmt.Println(out, err, uuid)
			{
				if err != nil {
					c.Abort()
					fmt.Print(out)
					model.Errorres(nil, "token错误", c)
				} else {
					if out == t {
						c.Next()
					} else {
						fmt.Print(out)
						c.Abort()
						model.Errorres(nil, "token错误", c)
					}
				}
			}
		} else {
			fmt.Println(err)
			fmt.Println("失效")
			c.Abort()
			model.Errorres(nil, "token错误", c)
		}
	}
}
