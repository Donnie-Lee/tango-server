package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type Configuration struct {
	DBConfig    *DBConfig    `yaml:"database"`
	RedisConfig *RedisConfig `yaml:"redis"`
}

var Config *Configuration

const Log_FILE_PATH string = "log.log"

func init() {
	// 加载配置
	err := loadConfig("application.yml")
	if err != nil {
		fmt.Println("Database Failed to load configuration")
		return
	}
}

func loadConfig(path string) error {
	result, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(result, &Config)
}
