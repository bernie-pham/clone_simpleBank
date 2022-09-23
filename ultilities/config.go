package ultilities

import (
	"github.com/spf13/viper"
)

type Config struct {
	DRIVER_SOURCE  string
	DRIVER_NAME    string
	SERVER_ADDRESS string
}

func LoadConfig(path string) (Config, error) {
	var config Config
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return config, err
	}

	err = viper.Unmarshal(&config)
	return config, err
}
