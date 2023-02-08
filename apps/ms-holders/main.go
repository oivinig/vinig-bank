package main

import (
	"github.com/oivinig/vinigbank/apps/msholders/adapter"
	"github.com/spf13/viper"
)

func main() {
	r := adapter.Router()

	//FIXME: read all configs from .env file
	viper.SetDefault("PGHOST", "0.0.0.0")
	viper.SetDefault("PGUSER", "postgres")
	viper.SetDefault("PGPASSWORD", "postgres")
	viper.SetDefault("HOSTNAME", "vinigbank-ms-holders")
	viper.SetDefault("HOST", "localhost")

	viper.BindEnv("PGHOST", "PGHOST")
	viper.BindEnv("PGUSER", "PGUSER")
	viper.BindEnv("PGPASSWORD", "PGPASSWORD")
	viper.BindEnv("HOSTNAME", "HOSTNAME")
	viper.BindEnv("HOST", "HOST")
	r.Run(":8001")
}
