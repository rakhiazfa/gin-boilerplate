package infrastructures

import (
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewPostgreSQLConnection() *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s port=%d dbname=%s user=%s password=%s sslmode=disable TimeZone=%s",
		viper.GetString("database.host"),
		viper.GetInt("database.port"),
		viper.GetString("database.name"),
		viper.GetString("database.username"),
		viper.GetString("database.password"),
		viper.GetString("application.timezone"),
	)

	gorm, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger:                 logger.Default.LogMode(logger.Info),
		SkipDefaultTransaction: true,
	})
	if err != nil {
		log.Fatal("Failed to initialize postgres connection: ", err)
	}

	db, err := gorm.DB()
	if err != nil {
		log.Fatal("Failed to get database connection: ", err)
	}

	db.SetMaxOpenConns(viper.GetInt("database.max_open_connections"))
	db.SetMaxIdleConns(viper.GetInt("database.max_idle_connections"))
	db.SetConnMaxLifetime(viper.GetDuration("database.max_connection_life_time") * time.Minute)
	db.SetConnMaxIdleTime(viper.GetDuration("database.max_connection_idle_time") * time.Minute)

	return gorm
}
