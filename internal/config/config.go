package config

/*
Package config
Подготавливает и загружает конфигурацию для служб сервиса.
Инстанциирует объекты для конфигурирования клиентов Postgres, Tarantool, Vault и.т.д
*/

import (
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/spf13/viper"
)

const (
	CONFIG_DIR       = "configs"
	CONFIG_PROD_FILE = "prod"
	CONFIG_DEV_FILE  = "dev"
)

type Config struct {
	Postgres Postgres
	Url      Url
	IsProd   bool

	Http struct {
		Port int64 `mapstructure:"port"`
	} `mapstructure:"http"`

	Cache struct {
		Ttl int64 `mapstructure:"ttl"`
	} `mapstructure:"cache"`

	Ctx struct {
		Ttl time.Duration `mapstructure:"ttl"`
	} `mapstructure:"ctx"`
}

type Postgres struct {
	Host      string
	Port      string
	CabinetDb string
	AdminDb   string
	User      string
	Password  string
	SSLMode   string
}

type Url struct {
	AuthGrpc string
}

type Telegram struct {
	BotToken string `vault:"telegram_bot_token"`
	ChatId   string `vault:"telegram_chat_id"`
}

func New(isProd bool) (*Config, error) {
	cfg := new(Config)

	viper.AddConfigPath(CONFIG_DIR)
	viper.SetConfigName(CONFIG_DEV_FILE)

	if isProd {
		viper.SetConfigName(CONFIG_PROD_FILE)
		cfg.IsProd = true
	}

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(cfg); err != nil {
		return nil, err
	}

	if err := envconfig.Process("postgres", &cfg.Postgres); err != nil {
		return nil, err
	}

	if err := envconfig.Process("url", &cfg.Url); err != nil {
		return nil, err
	}

	return cfg, nil
}
