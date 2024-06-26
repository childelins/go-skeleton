package env

import "github.com/childelins/go-skeleton/pkg/config"

func IsLocal() bool {
	return config.Get("app.env") == "local"
}

func IsTesting() bool {
	return config.Get("app.env") == "testing"
}

func IsProduction() bool {
	return config.Get("app.env") == "production"
}
