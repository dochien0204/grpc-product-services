package config

import "github.com/spf13/viper"

type DBConfig struct {
	Host         string `mapstructure:"DB_HOST"`
	User         string `mapstructure:"DB_USER"`
	Password     string `mapstructure:"DB_PASSWORD"`
	DatabaseName string `mapstructure:"DB_NAME"`
	Port         string `mapstructure:"DB_PORT"`
}

func LoadConfig() (c DBConfig, err error) {
	viper.AddConfigPath("./pkg/config/envs")
	viper.SetConfigName("db")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&c)

	return

}
