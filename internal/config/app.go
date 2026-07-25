package config

const (
	APP_NAME    = getEnv("APP_NAME", "User REST API")
	APP_URL     = getEnv("APP_URL", "http://localhost")
	APP_PORT    = getEnv("APP_PORT", "8080")
	APP_VERSION = "v1.0"
)
