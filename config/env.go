package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	PublicHost          string
	Port                string
	BucketPath          string
	BufferSize          int
	BucketNameMaxLength int
}

var Envs = initConfig()

func initConfig() Config {
	godotenv.Load()
	return Config{
		PublicHost:          getEnv("PUBLIC_HOST", "localhost"),
		Port:                getEnv("PORT", "8080"),
		BucketPath:          getEnv("BUCKET_PATH", "buckets/"),
		BufferSize:          getEnvInt("BUFFER_SIZE", 1024),
		BucketNameMaxLength: getEnvInt("BUCKET_NAME_MAX_LENGTH", 100),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getEnvInt(key string, fallback int) int {
	if value, ok := os.LookupEnv(key); ok {
		v, err := strconv.Atoi(value)
		if err != nil {
			log.Fatalf("environment variable %s could not be converted to type int \n error: %v", key, err)
		}
		return v
	}
	return fallback
}
