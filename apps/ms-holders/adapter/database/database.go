package database

import (
	"fmt"
	"github.com/oivinig/vinigbank/apps/msholders/adapter/database/model"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (db *gorm.DB) {
	host := viper.Get("PGHOST")
	user := viper.Get("PGUSER")
	pass := viper.Get("PGPASSWORD")
	dsn := fmt.Sprintf("host=%v user=%v password=%v", host, user, pass)

	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(&model.Holder{})
	sqlDB, err := db.DB()
	if err != nil {
		panic(err.Error())
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	return db
}
