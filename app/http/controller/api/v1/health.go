package v1

import (
	"github.com/childelins/go-skeleton/pkg/resp"
	"github.com/gin-gonic/gin"
)

type HealthController struct {
	BaseController
}

func NewHealthController() *HealthController {
	return &HealthController{}
}

func (h *HealthController) Check(c *gin.Context) {
	// panic("out of memory")
	// resp.Success(c)
	// resp.Error(c, "error")
	// resp.WithErrCode(c, errcode.ERROR)
	// resp.WithErrCode(c, errcode.New(10001, "not found"))
	resp.Success(c, "ok")
}
