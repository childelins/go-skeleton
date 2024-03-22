package bootstrap

import (
	"github.com/childelins/go-skeleton/pkg/config"
	"github.com/childelins/go-skeleton/pkg/container"
	"github.com/childelins/go-skeleton/pkg/log"
	"github.com/childelins/go-skeleton/pkg/routes"
	"go.uber.org/zap"
)

type App struct {
	c *container.Container
}

func NewApp(c *container.Container) *App {
	return &App{c: c}
}

func (a *App) Run() {
	r := routes.New(a.c)

	// 运行服务器
	if err := r.Run(":" + config.Get("app.port")); err != nil {
		log.Fatal("Unable to start server", zap.Error(err))
	}
}
