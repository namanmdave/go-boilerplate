package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type DBConfig struct {
	DBHost     string
	DBPort     int
	DBUser     string
	DBPassword string
	DBName     string
}

func (c *DBConfig) GetDSN() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		c.DBHost, c.DBPort, c.DBUser, c.DBPassword, c.DBName)
}

var dbConfig *DBConfig

func GetDBConfig() *DBConfig {
	if dbConfig != nil {
		return dbConfig
	}

	dbConfig = &DBConfig{
		DBHost:     getEnvString("DB_HOST", "localhost"),
		DBPort:     getEnvInt("DB_PORT", 5432),
		DBName:     getEnvString("DB_NAME", "boilerplate_db"),
		DBPassword: getEnvString("DB_PASSWORD", "password123"),
		DBUser:     getEnvString("DB_USER", "boilerplate"),
	}

	return dbConfig
}

type RedisConfig struct {
	Addr     string
	Password string
	DB       int
}

var redisConfig *RedisConfig

func GetRedisConfig() *RedisConfig {
	if redisConfig != nil {
		return redisConfig
	}

	host := getEnvString("REDIS_HOST", "localhost")
	port := getEnvString("REDIS_PORT", "6379")

	redisConfig = &RedisConfig{
		Addr:     host + ":" + port,
		Password: getEnvString("REDIS_PASSWORD", ""),
		DB:       getEnvInt("REDIS_DB", 0),
	}

	return redisConfig
}

type ServerConfig struct {
	HTTPAddr string
}

var serverConfig *ServerConfig

func GetServerConfig() *ServerConfig {
	if serverConfig != nil {
		return serverConfig
	}

	port := getEnvString("HTTP_PORT", "8080")
	serverConfig = &ServerConfig{
		HTTPAddr: ":" + port,
	}

	return serverConfig
}

// Helper functions to get environment variables with defaults
func getEnvString(key, defaultValue string) string {
	viper.SetDefault(key, defaultValue)
	return viper.GetString(key)
}

func getEnvInt(key string, defaultValue int) int {
	viper.SetDefault(key, defaultValue)
	return viper.GetInt(key)
}
