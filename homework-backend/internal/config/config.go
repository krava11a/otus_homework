package config

import (
	"flag"
	"homework-backend/internal/models"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env                string     `yaml:"env" env-default:"local"`
	StoragePath        string     `yaml:"storage_path" env-required:"true"`
	StoragePathForRead string     `yaml:"storage_path_for_read" env-required:"false"`
	CachePath          string     `yaml:"cache_path" env-required:"false"`
	QueuePath          string     `yaml:"rqueue_path" env-required:"false"`
	GRPC               GRPCConfig `yaml:"grpc"`
	MigrationsPath     string
	TokenTTL           time.Duration `yaml:"token_ttl" env-default:"1h"`
	App                models.App    `yaml:"app"`
}

type GRPCConfig struct {
	Port    uint          `yaml:"port"`
	WebPort uint          `yaml:"web_port"`
	WsPort  uint          `yaml:"ws_port"`
	Timeout time.Duration `yaml:"timeout"`
}

func MustLoad() *Config {
	configPath := fetchConfigPath()
	if configPath == "" {
		panic("config path is empty")
	}

	// check if file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		panic("config file does not exist: " + configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		panic("config path is empty: " + err.Error())
	}

	return &cfg
}

// fetchConfigPath fetches config path from command line flag or environment variable.
// Priority: flag > env > default.
// Default value is empty string.
func fetchConfigPath() string {
	var res string

	flag.StringVar(&res, "config", "", "path to config file")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}

	return res
}
