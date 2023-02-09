package main

import (
	"fmt"

	"github.com/oivinig/vinigbank/apps/msholders/adapter"
	"github.com/spf13/viper"
)

func main() {
	r := adapter.Router()
	readConfig()
	port := viper.GetString("PORT")

	r.Run(fmt.Sprintf(":%v", port))
}

func readConfig() {
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}
