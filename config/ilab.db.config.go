package config

import (
	"../structs"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

//DB Credential, Change if Necessary
//const username = "ilab-demo"
//const password = "demodemodemodemodemodemodemo"
//const dbName = "ilab-db-demo"
//var (
//	username = os.Getenv("DB_CLIENT_USER")
//	password = os.Getenv("DB_CLIENT_USER_PASSWORD")
//	dbName = os.Getenv("DB_ILAB_NAME")
//	dbName2 = os.Getenv("DB_CLIENT_NAME")
//)

const username = "user1"
const password = ""
const dbName = "infotech.ilab"
const dbName2 = "infotech.ilab.auth.client"

//TODO: Documentation
func DBInit() *gorm.DB {
	//dsn := fmt.Sprintf("ocs_main:5a64ffaf5121a286a45d97bed73fc7d41337d0af442638279d6b067bb329afa9@(10.10.11.254:3306)/%s?charset=utf8mb4&parseTime=true&loc=Local", dbName)
	dsn := fmt.Sprintf("user1:@(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=true&loc=Local", dbName)

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

func ClientDBInit() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@(127.0.0.1)/%s?charset=utf8mb4&parseTime=true&loc=Local", username, password, dbName2)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic(err)
	}

	_ = db.AutoMigrate(
		structs.Client{},
	)

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	sqlDB.SetMaxIdleConns(0)
	return db
}
