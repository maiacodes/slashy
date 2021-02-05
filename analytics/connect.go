package analytics

import (
	"gorm.io/driver/bigquery"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
)

var (
	G *gorm.DB
)

func Connect() {
	db, err := gorm.Open(bigquery.Open("bigquery://slashy/track"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
	})
	if err != nil {
		panic(err)
	}

	// Make sure we have a complex_records table
	if os.Getenv("development") == "true" {
		err = db.AutoMigrate(&event{})
		if err != nil {
			panic(err)
		}
	}

	G = db
}
