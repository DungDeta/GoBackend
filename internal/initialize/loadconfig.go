package initialize

import (
	"fmt"

	viper2 "github.com/spf13/viper"
	"myproject/global"
)

func LoadConfig() {
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

	if err := viper.Unmarshal(&global.Config); err != nil {
		fmt.Println(err)
	}

	fmt.Println("Server Port: ", global.Config.Server.Port)
}
