package model

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

const (
	ERROR   = 500
	SUCCESS = 200
)

func Result(code int, data interface{}, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

// Successres 传入空接口，信息
func Successres(data interface{}, message string, c *gin.Context) {
	Result(SUCCESS, data, message, c)
}

//
func Errorres(data interface{}, message string, c *gin.Context) {
	Result(ERROR, data, message, c)
}
