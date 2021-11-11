package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// マッピング用の構造体
type Config struct {
	User User `yaml:user`
}

type User struct {
	Name string `yaml:"name"`
}

func Load() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config/")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("設定ファイル読み込みエラー: %s \n", err)
	}

	var cfg Config

	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, fmt.Errorf("unmarshal error: %s \n", err)
	}

	return &cfg, nil
}
