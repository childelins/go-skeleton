//go:build wireinject
// +build wireinject

package main

import (
	api "github.com/childelins/go-skeleton/app/http/controller/api/v1"
	"github.com/childelins/go-skeleton/app/repository"
	"github.com/childelins/go-skeleton/app/service"
	"github.com/childelins/go-skeleton/bootstrap"
	"github.com/childelins/go-skeleton/pkg/container"
	"github.com/google/wire"
	"gorm.io/gorm"
)

func BuildApp(db *gorm.DB) *bootstrap.App {
	wire.Build(
		bootstrap.NewApp,
		container.New,
		api.ControllerSet,
		service.ServiceSet,
		repository.RepositorySet,
	)

	return &bootstrap.App{}
}
