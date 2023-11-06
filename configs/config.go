package configs

import (	
	"github.com/spf13/viper"
)

type config struct {
	DBHost			string `mapstructure:"DB_HOST"`
	DBUser			string `mapstructure:"DB_USER"`
	DBPassword		string `mapstructure:"DB_PASSWORD"`
	DBPort			string `mapstructure:"DB_PORT"`
	DBName			string `mapstructure:"DB_NAME"`
	Port			string `mapstructure:"SERVER_PORT"`
	JwtSecret		string `mapstructure:"JWT_SECRET"`
}

func InitConfig(path string) (cfg config, err error) {
	
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&cfg)
	return
}