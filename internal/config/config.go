package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	KafkaBrokers   string `envconfig:"KAFKA_BROKERS" required:"true"`
	KafkaTopic     string `envconfig:"KAFKA_TOPIC" required:"true"`
	PostgresqlDSN  string `envconfig:"POSTGRESQL_DSN" required:"true"`
	HTTPServerPort string `envconfig:"HTTP_SERVER_PORT" default:"8080"`
}

func LoadConfig() (*Config, error) {
	var cfg Config
	err := envconfig.Process("", &cfg)
	return &cfg, err

}
