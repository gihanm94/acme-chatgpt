package config

import (
	"github.com/spf13/viper"
)

func LoadConfig() *viper.Viper {
	v := viper.New()
	v.AutomaticEnv()
	v.SetDefault("port", 80)
	v.SetDefault("lark_base_url", "https://open.larksuite.com")
	v.SetDefault("bot_name", "chatgpt")

	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath("./../..")
	v.AddConfigPath("./../")
	v.AddConfigPath("./")
	_ = v.ReadInConfig()

	return v
}