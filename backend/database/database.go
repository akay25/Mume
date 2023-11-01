package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectToDB(dbPath string) (db *gorm.DB, err error) {
	DB, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	return DB, err
}
