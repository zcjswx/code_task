//go:build test
// +build test

package app

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	con, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db = con
	db.AutoMigrate(&EdgeEntity{}, &GraphEntity{}, &NodeEntity{})
}
