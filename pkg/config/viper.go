package config

import "github.com/spf13/viper"

func InitViper(path, filename, ext string) {
	viper.AddConfigPath(path)
	viper.SetConfigName(filename)
	viper.SetConfigType(ext)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
