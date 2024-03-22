package routes

import (
	"github.com/childelins/go-skeleton/pkg/container"
	"github.com/gin-gonic/gin"
)

// RegisterAPIRoutes 注册 API 相关路由
func RegisterAPIRoutes(c *container.Container, r *gin.RouterGroup) {
	r.GET("/health", c.HealthController.Check)

	r.GET("/user/create", c.UserController.Create)
	r.GET("/user/list", c.UserController.List)
}
