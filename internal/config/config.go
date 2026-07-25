package config

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

// Application
var (
	APP_NAME     string
	APP_ENV      string
	APP_PORT     string
	APP_DEBUG    bool
	APP_VERSION  string
	APP_TIMEZONE string
	APP_URL      string
)

// API
var (
	API_PREFIX  string
	API_VERSION string
)

// Database
var (
	DB_CONNECTION         string
	DB_HOST               string
	DB_PORT               string
	DB_NAME               string
	DB_USER               string
	DB_PASSWORD           string
	DB_MAX_OPEN_CONNS     int
	DB_MAX_IDLE_CONNS     int
	DB_CONN_MAX_LIFETIME  time.Duration
	DB_CONN_MAX_IDLE_TIME time.Duration
)

// SMTP
var (
	SMTP_HOST     string
	SMTP_PORT     int
	SMTP_USERNAME string
	SMTP_PASSWORD string
	SMTP_FROM     string
)

// Redis
var (
	REDIS_HOST     string
	REDIS_PORT     int
	REDIS_PASSWORD string
	REDIS_DB       int
)

// Logging
var (
	LOG_LEVEL string
)

// File Uploads
var (
	UPLOAD_DIR      string
	MAX_UPLOAD_SIZE string
)

// Pagination
var (
	DEFAULT_PAGE_SIZE int
	MAX_PAGE_SIZE     int
)

// Rate Limiting
var (
	RATE_LIMIT_REQUESTS int
	RATE_LIMIT_DURATION time.Duration
)

// JWT Authentication
var (
	JWT_SECRET     string
	JWT_EXPIRATION time.Duration
	JWT_ISSUER     string
)

// CORS
var (
	CORS_ALLOWED_ORIGINS string
	CORS_ALLOWED_METHODS string
	CORS_ALLOWED_HEADERS string
)

func init() {
	// Load .env file
	err := godotenv.Load()
	if err != nil && os.Getenv("APP_ENV") != "production" {
		log.Println("Warning: .env file not found")
	}

	// Application
	APP_NAME = getEnv("APP_NAME", "User REST API")
	APP_ENV = getEnv("APP_ENV", "development")
	APP_PORT = getEnv("APP_PORT", "8080")
	APP_DEBUG = getEnvBool("APP_DEBUG", false)
	APP_VERSION = getEnv("APP_VERSION", "v1.0")
	APP_TIMEZONE = getEnv("APP_TIMEZONE", "UTC")
	APP_URL = getEnv("APP_URL", "http://localhost")

	// API
	API_PREFIX = getEnv("API_PREFIX", "/api")
	API_VERSION = getEnv("API_VERSION", "v1")

	// Database
	DB_CONNECTION = getEnv("DB_CONNECTION", "mysql")
	DB_HOST = getEnv("DB_HOST", "127.0.0.1")
	DB_PORT = getEnv("DB_PORT", "3306")
	DB_NAME = getEnv("DB_NAME", "")
	DB_USER = getEnv("DB_USER", "")
	DB_PASSWORD = getEnv("DB_PASSWORD", "")
	DB_MAX_OPEN_CONNS = getEnvInt("DB_MAX_OPEN_CONNS", 10)
	DB_MAX_IDLE_CONNS = getEnvInt("DB_MAX_IDLE_CONNS", 10)
	DB_CONN_MAX_LIFETIME = getEnvDuration("DB_CONN_MAX_LIFETIME", 5*time.Minute)
	DB_CONN_MAX_IDLE_TIME = getEnvDuration("DB_CONN_MAX_IDLE_TIME", 2*time.Minute)

	// SMTP
	SMTP_HOST = getEnv("SMTP_HOST", "")
	SMTP_PORT = getEnvInt("SMTP_PORT", 587)
	SMTP_USERNAME = getEnv("SMTP_USERNAME", "")
	SMTP_PASSWORD = getEnv("SMTP_PASSWORD", "")
	SMTP_FROM = getEnv("SMTP_FROM", "")

	// Redis
	REDIS_HOST = getEnv("REDIS_HOST", "127.0.0.1")
	REDIS_PORT = getEnvInt("REDIS_PORT", 6379)
	REDIS_PASSWORD = getEnv("REDIS_PASSWORD", "")
	REDIS_DB = getEnvInt("REDIS_DB", 0)

	// Logging
	LOG_LEVEL = getEnv("LOG_LEVEL", "info")

	// File Uploads
	UPLOAD_DIR = getEnv("UPLOAD_DIR", "./storage/uploads")
	MAX_UPLOAD_SIZE = getEnv("MAX_UPLOAD_SIZE", "10MB")

	// Pagination
	DEFAULT_PAGE_SIZE = getEnvInt("DEFAULT_PAGE_SIZE", 10)
	MAX_PAGE_SIZE = getEnvInt("MAX_PAGE_SIZE", 100)

	// Rate Limiting
	RATE_LIMIT_REQUESTS = getEnvInt("RATE_LIMIT_REQUESTS", 100)
	RATE_LIMIT_DURATION = getEnvDuration("RATE_LIMIT_DURATION", 1*time.Minute)

	// JWT Authentication
	JWT_SECRET = getEnv("JWT_SECRET", "")
	JWT_EXPIRATION = getEnvDuration("JWT_EXPIRATION", 24*time.Hour)
	JWT_ISSUER = getEnv("JWT_ISSUER", "UserRestAPIGo")

	// CORS
	CORS_ALLOWED_ORIGINS = getEnv("CORS_ALLOWED_ORIGINS", "http://localhost:3000")
	CORS_ALLOWED_METHODS = getEnv("CORS_ALLOWED_METHODS", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
	CORS_ALLOWED_HEADERS = getEnv("CORS_ALLOWED_HEADERS", "Content-Type,Authorization")

	// Validate required variables
	validateRequired("DB_NAME", DB_NAME)
	validateRequired("DB_USER", DB_USER)
	// validateRequired("JWT_SECRET", JWT_SECRET)
}

// Helper functions
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvBool(key string, defaultValue bool) bool {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	boolVal, err := strconv.ParseBool(value)
	if err != nil {
		log.Printf("Invalid boolean value for %s: %s, using default: %v", key, value, defaultValue)
		return defaultValue
	}
	return boolVal
}

func getEnvInt(key string, defaultValue int) int {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	intVal, err := strconv.Atoi(value)
	if err != nil {
		log.Printf("Invalid integer value for %s: %s, using default: %d", key, value, defaultValue)
		return defaultValue
	}
	return intVal
}

func getEnvDuration(key string, defaultValue time.Duration) time.Duration {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	duration, err := time.ParseDuration(value)
	if err != nil {
		log.Printf("Invalid duration value for %s: %s, using default: %v", key, value, defaultValue)
		return defaultValue
	}
	return duration
}

func validateRequired(key, value string) {
	if value == "" {
		log.Fatalf("Environment variable %s is required", key)
	}
}
