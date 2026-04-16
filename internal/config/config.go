package config

import (
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	JWT      JwtConfig
	AWS      AwsConfig
	Upload   UploadConfig
}

type ServerConfig struct {
	Port     string
	GinDebug string
}

type DatabaseConfig struct {
	DBhost     string
	DBport     int
	DBuser     string
	DBpassword string
	DBname     string
	DBsslmode  string
}

type JwtConfig struct {
	JwtSecret    string
	JwtExp       time.Duration
	RefreshToken time.Duration
}

type AwsConfig struct {
	Region          string
	AccessKeyID     string
	SecretAccessKey string
	S3Bucket        string
	S3Endpoint      string
}

type UploadConfig struct {
	Region      string
	MaxFileSize int64
}

func Load() (*Config, error) {
	_ = godotenv.Load()

	jwtExpired, _ := time.ParseDuration(getEnv("JWT_EXPIRATION", "24h"))
	refreshTokenExpired, _ := time.ParseDuration(getEnv("JWT_REFRESH_EXPIRATION", "168h"))
	maxUploadSize, _ := strconv.ParseInt(getEnv("MAX_UPLOAD_SIZE", "10485760"), 10, 64) // 10MB default

	dbPort, _ := strconv.Atoi(getEnv("DB_PORT", "5432"))

	config := &Config{
		Server: ServerConfig{
			Port:     getEnv("APP_PORT", "8080"),
			GinDebug: getEnv("APP_GIN_MODE", "Debug"),
		},
		Database: DatabaseConfig{
			DBhost:     getEnv("DB_HOST", "localhost"),
			DBport:     dbPort,
			DBuser:     getEnv("DB_USER", "ucommerce"),
			DBpassword: getEnv("DB_PASSWORD", "admin1234"),
			DBname:     getEnv("DB_NAME", "ecommerce-shop"),
			DBsslmode:  getEnv("DB_SSL_MODE", "disable"),
		},
		JWT: JwtConfig{
			JwtSecret:    getEnv("JWT_SECRET", "your-secret-key-here"),
			JwtExp:       jwtExpired,
			RefreshToken: refreshTokenExpired,
		},
		AWS: AwsConfig{
			Region:          getEnv("AWS_DEFAULT_REGION", "us-east-1"),
			AccessKeyID:     getEnv("AWS_ACCESS_KEY_ID", "test"),
			SecretAccessKey: getEnv("AWS_SECRET_ACCESS_KEY", "test"),
			S3Bucket:        getEnv("S3_BUCKET", "ecommerce-bucket"),
			S3Endpoint:      getEnv("AWS_ENDPOINT_URL", "http://localhost:4566"),
		},
		Upload: UploadConfig{
			Region:      getEnv("AWS_DEFAULT_REGION", "us-east-1"),
			MaxFileSize: maxUploadSize,
		},
	}

	return config, nil
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)

	if value == "" {
		return defaultValue
	}

	return value
}
