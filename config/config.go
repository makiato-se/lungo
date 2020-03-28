package config

import "github.com/caarlos0/env"

type Config struct {
	Port          string `env:"PORT" envDefault:"8080"`
	MongoEndpoint string `env:"MONGODB_ENDPOINT" envDefault:"mongodb://localhost:27017"`
}

func ParseConfig() *Config {
	conf := &Config{}
	err := env.Parse(conf)
	if err != nil {
		panic(err)
	}

	return conf
}
