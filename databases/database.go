package internal

import (
	"fmt"
	"log"
	"os"
	"think-intern-2023/logs"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func CreateDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v TimeZone=%v",
		viper.GetString("DB_HOST"),
		viper.GetString("DB_USERNAME"),
		viper.GetString("DB_PASSWORD"),
		viper.GetString("DB_DATABASE"),
		viper.GetInt("DB_PORT"),
		viper.GetString("DB_SSLMODE"),
		viper.GetString("DB_TIMEZONE"))

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,          // Don't include params in the SQL log
			Colorful:                  true,          // Disable color
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
		DryRun: false,
	})
	if err != nil {
		logs.Error(err)
	}

	table := fmt.Sprintf("set search_path=%v", viper.GetString("DB_TABLENAME"))
	db.Exec(table)

	return db, nil
}

func AutoMigrate(am *gorm.DB) {
	am.Debug().AutoMigrate()
}
