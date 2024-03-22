package routes

import (
	"net/http"
	"strings"

	"github.com/childelins/go-skeleton/app/http/middleware"
	"github.com/childelins/go-skeleton/pkg/config"
	"github.com/childelins/go-skeleton/pkg/container"
	"github.com/childelins/go-skeleton/pkg/errcode"
	"github.com/childelins/go-skeleton/pkg/resp"
	"github.com/childelins/go-skeleton/routes"
	"github.com/gin-gonic/gin"
)

type Router struct {
	r *gin.Engine

	c *container.Container
}

func New(c *container.Container) *Router {
	// 设置 gin 的运行模式，支持 debug, release, test
	// release 会屏蔽调试信息，官方建议生产环境中使用
	// 非 release 模式 gin 终端打印太多信息，干扰到我们程序中的 Log
	// 故此设置为 release，有特殊情况手动改为 debug 即可
	gin.SetMode(gin.ReleaseMode)

	// gin 实例
	r := gin.New()

	return &Router{r: r, c: c}
}

func (r *Router) Run(addr ...string) error {
	// 初始化路由绑定
	r.SetupRoute()

	return r.r.Run(addr...)
}

// SetupRoute 路由初始化
func (r *Router) SetupRoute() {
	// 注册全局中间件
	r.registerGlobalMiddleWare()

	// 注册 API 路由
	r.registerAPIRoutes()

	// 配置 404 路由
	r.setup404Handler()
}

func (r *Router) registerGlobalMiddleWare() {
	r.r.Use(
		middleware.Recovery(),
	)
}

func (r *Router) registerAPIRoutes() {
	// 我们所有的 v1 版本的路由都将存放到这里
	var v1 *gin.RouterGroup
	if len(config.Get("app.api_domain")) == 0 {
		v1 = r.r.Group("/api/v1")
	} else {
		v1 = r.r.Group("/v1")
	}

	routes.RegisterAPIRoutes(r.c, v1)
}

func (r *Router) setup404Handler() {
	// 处理 404 请求
	r.r.NoRoute(func(c *gin.Context) {
		// 获取标头信息的 Accept 信息
		acceptString := c.Request.Header.Get("Accept")
		if strings.Contains(acceptString, "text/html") {
			// 如果是 HTML 的话
			c.String(http.StatusNotFound, "404 not found")
		} else {
			// 默认返回 JSON
			resp.WithErrCode(c, errcode.NOT_FOUND)
		}
	})
}
