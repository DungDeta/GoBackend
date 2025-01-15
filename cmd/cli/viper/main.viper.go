package main

import (
	"fmt"

	viper2 "github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`
	Security struct {
		Jwt struct {
			Key string `mapstructure:"key"`
		}
	} `mapstructure:"security"`
	Databases []struct {
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Host     string `mapstructure:"host"`
		DbName   string `mapstructure:"dbName"`
	} `mapstructure:"databases"`
}

func main() {
	viper := viper2.New()
	viper.AddConfigPath("./configs") // Path to configs
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig() // Đọc file config
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	fmt.Println("Server Port: ", viper.Get("server.port"))
	fmt.Println("Security Key: ", viper.Get("security.jwt.key"))

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Println(err)
	}

	fmt.Println("Server Port: ", config.Server.Port)
	for _, db := range config.Databases {
		fmt.Println("User: ", db.User)
	}
}
