package db

import (
	"os"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var err error

func Db() *gorm.DB {

	if db == nil {
		db, err = gorm.Open(
			"mysql",
			os.Getenv("DBURL"))
		if err != nil {
			panic("failed to connect database")
		}
		db.SingularTable(true)
		db.DB().SetMaxIdleConns(0)
		db.DB().SetMaxOpenConns(5)
		db.DB().SetConnMaxLifetime(86400)
	}

	// defer db.Close()
	return db
}
