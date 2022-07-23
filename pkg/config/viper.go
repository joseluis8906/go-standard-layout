package config

import "github.com/spf13/viper"

// Load ...
func Load(path, filename, ext string) {
	viper.AddConfigPath(path)
	viper.SetConfigName(filename)
	viper.SetConfigType(ext)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
