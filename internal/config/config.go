package config

import (
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
)

type Config struct {
	Host      string    `yaml:"host" env-default:"127.0.0.1"`
	Port      string    `yaml:"port" env-default:"3000"`
	JwtSecret string    `yaml:"jwt_secret" env-default:"secret"`
	Memcached Memcached `yaml:"memcached"`
	Mongo     Mongo     `yaml:"mongo"`
}

type Mongo struct {
	Database string `yaml:"database" env-default:"url-shortener"`
	URI      string `yaml:"uri"`
	Username string `yaml:"username" env-default:"admin"`
	Password string `yaml:"password" env-default:"admin"`
}

type Memcached struct {
	URI     string `yaml:"uri" env-default:"localhost:11211"`
	MaxSize int    `yaml:"max_size" env-default:"1000"`
}

var configPath string

func parseConfigPath() {
	flag.StringVar(&configPath, "config", configPath, "config file path")
	flag.Parse()

	if configPath == "" {
		log.Fatalf("config file path not specified or unreachable")
	}
}

func NewConfig() *Config {
	var config Config

	parseConfigPath()
	if err := cleanenv.ReadConfig(configPath, &config); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	return &config
}
