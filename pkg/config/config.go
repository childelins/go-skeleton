package config

import (
	"os"

	"github.com/spf13/cast"
	viperlib "github.com/spf13/viper"
)

// viper 库实例
var viper *viperlib.Viper

// ConfigFunc 动态加载配置信息
type ConfigFunc func() map[string]interface{}

// ConfigFuncs 先加载到此数组，loadConfig 再动态生成配置信息
var ConfigFuncs map[string]ConfigFunc

func init() {
	// 1. 初始化 Viper 库
	viper = viperlib.New()
	// 2. 配置类型，支持 "json", "toml", "yaml", "yml", "properties",
	//             "props", "prop", "env", "dotenv"
	viper.SetConfigType("env")
	// 3. 环境变量配置文件查找的路径，相对于 main.go
	viper.AddConfigPath(".")
	// 4. 设置环境变量前缀，用以区分 Go 的系统环境变量
	viper.SetEnvPrefix("appenv")
	// 5. 读取环境变量（支持 flags）
	viper.AutomaticEnv()

	ConfigFuncs = make(map[string]ConfigFunc)
}

// InitConfig 初始化配置信息，完成对环境变量以及 config 信息的加载
func InitConfig(env string) {
	// 1. 加载环境变量
	loadEnv(env)
	// 2. 注册配置信息
	loadConfig()
}

func loadEnv(envSuffix string) {
	// 默认加载 .env 文件，如果有传参 --env=name 的话，加载 .env.name 文件
	envPath := ".env"

	if len(envSuffix) > 0 {
		filepath := ".env." + envSuffix
		if _, err := os.Stat(filepath); err == nil {
			// 如 .env.testing 或 .env.stage
			envPath = filepath
		}
	}

	// 加载 env
	viper.SetConfigName(envPath)
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	// // 监控 .env 文件，变更时重新加载
	// viper.WatchConfig()
}

func loadConfig() {
	for name, fn := range ConfigFuncs {
		// 设置配置信息
		viper.Set(name, fn())
	}
}

// Env 读取环境变量，支持默认值
func Env(envName string, defaultValue ...interface{}) interface{} {
	return internalGet(envName, defaultValue...)
}

// Add 新增配置项
func Add(name string, configFunc ConfigFunc) {
	ConfigFuncs[name] = configFunc
}

// Get 获取配置项
// 第一个参数 key 允许使用点式获取，如：app.name
// 第二个参数允许传参默认值
func Get(key string, defaultValue ...interface{}) string {
	return GetString(key, defaultValue...)
}

// GetString 获取 String 类型的配置信息
func GetString(key string, defaultValue ...interface{}) string {
	return cast.ToString(internalGet(key, defaultValue...))
}

// GetInt 获取 Int 类型的配置信息
func GetInt(key string, defaultValue ...interface{}) int {
	return cast.ToInt(internalGet(key, defaultValue...))
}

// GetFloat64 获取 float64 类型的配置信息
func GetFloat64(key string, defaultValue ...interface{}) float64 {
	return cast.ToFloat64(internalGet(key, defaultValue...))
}

// GetInt64 获取 Int64 类型的配置信息
func GetInt64(key string, defaultValue ...interface{}) int64 {
	return cast.ToInt64(internalGet(key, defaultValue...))
}

// GetUint 获取 Uint 类型的配置信息
func GetUint(key string, defaultValue ...interface{}) uint {
	return cast.ToUint(internalGet(key, defaultValue...))
}

// GetBool 获取 Bool 类型的配置信息
func GetBool(key string, defaultValue ...interface{}) bool {
	return cast.ToBool(internalGet(key, defaultValue...))
}

// GetStringMapString 获取结构数据
func GetStringMapString(key string) map[string]string {
	return viper.GetStringMapString(key)
}

func internalGet(key string, defaultValue ...interface{}) interface{} {
	// config 或者环境变量不存在的情况
	if !viper.IsSet(key) {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return nil
	}

	return viper.Get(key)
}
