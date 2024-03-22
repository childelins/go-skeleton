package container

import api "github.com/childelins/go-skeleton/app/http/controller/api/v1"

type Container struct {
	HealthController *api.HealthController
	UserController   *api.UserController
}

func New(healthController *api.HealthController, userController *api.UserController) *Container {
	return &Container{
		HealthController: healthController,
		UserController:   userController,
	}
}
