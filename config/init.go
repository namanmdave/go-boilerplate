package config

import (
	"github.com/spf13/viper"
)

func Init() {
	// Initialize viper configuration from environment variables
	viper.AutomaticEnv()
	viper.SetDefault("server.http_addr", ":8080")
	viper.SetDefault("db.port", 5432)
	viper.SetDefault("redis.db", 0)

	// Initialize config
	_ = GetDBConfig()
	_ = GetRedisConfig()
	_ = GetServerConfig()
}
