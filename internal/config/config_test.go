package config

import (
	"testing"
)

func TestNewConfig(t *testing.T) {
	configPath = "../../config/config.yaml"
	config := NewConfig()
	t.Log(config)
}

func TestParseConfigPath(t *testing.T) {
	configPath = "../../config/config.yaml"
	parseConfigPath()
	t.Log(configPath)
}
