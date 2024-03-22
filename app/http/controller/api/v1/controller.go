package v1

import "github.com/google/wire"

var ControllerSet = wire.NewSet(
	NewHealthController,
	NewUserController,
)
