package models

import (
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func ConnectDatabase() {
  database, error := gorm.Open("sqlite3", "db.db")

  if error != nil {
    panic("Failed to connect to DB!")
  }

  database.AutoMigrate(&Movie{})

  DB = database
}
