package test

import (
	"testing"

	"github.com/childelins/go-skeleton/pkg/config"
)

func TestConfig(t *testing.T) {
	config.Add("test", func() map[string]interface{} {
		return map[string]interface{}{
			"name": "test",
		}
	})

	config.InitConfig("example")

	t.Log(config.GetString("test.name"))
	t.Log(config.Env("APP_ENV"))
}
