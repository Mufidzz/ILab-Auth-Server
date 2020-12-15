package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

//DB Credential, Change if Necessary
const username = "ilab-demo"
const password = "demodemodemodemodemodemodemo"
const dbName = "ilab-db-demo"

//TODO: Documentation
func DBInit() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@(127.0.0.1)/%s?charset=utf8mb4&parseTime=true&loc=Local", username, password, dbName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	sqlDB.SetMaxIdleConns(0)
	return db
}
