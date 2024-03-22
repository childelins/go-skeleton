package resp

import (
	"net/http"

	"github.com/childelins/go-skeleton/pkg/errcode"
	"github.com/gin-gonic/gin"
)

type response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func Success(c *gin.Context, data ...interface{}) {
	resp := response{
		Code: errcode.SUCCESS.Code(),
		Msg:  errcode.SUCCESS.Msg(),
	}
	if len(data) > 0 {
		resp.Data = data[0]
	}

	c.JSON(http.StatusOK, resp)
}

func Error(c *gin.Context, msg string) {
	c.AbortWithStatusJSON(http.StatusInternalServerError, response{
		Code: errcode.ERROR.Code(),
		Msg:  msg,
	})
}

func WithErrCode(c *gin.Context, err *errcode.ErrCode) {
	c.AbortWithStatusJSON(http.StatusInternalServerError, response{
		Code: err.Code(),
		Msg:  err.Msg(),
	})
}
