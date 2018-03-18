package database

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB, _ = gorm.Open("sqlite3", "test.db")
