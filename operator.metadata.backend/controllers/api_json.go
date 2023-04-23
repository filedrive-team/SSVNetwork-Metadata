package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func JSONOk(c *gin.Context, data interface{}) {
	if data == nil {
		data = struct{}{}
	}
	c.JSON(http.StatusOK, CommonOkResp{
		Code: http.StatusOK,
		Msg:  "OK",
		Data: data,
	})
}

func JSONError(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, CommonErrorResp{
		Code: http.StatusBadRequest,
		Msg:  msg,
	})
}
