package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"todo/backend/Config"
	"todo/backend/Models"
	"todo/backend/Routes"
)

var err error

func main() {
	godotenv.Load()

	Config.DB, err = gorm.Open(mysql.New(mysql.Config{
		DSN: Config.DbURL(Config.BuildDBConfig()), // data source name, refer https://github.com/go-sql-driver/mysql#dsn-data-source-name
	}), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	if err != nil {
		fmt.Println("Status:", err)
	}
	Config.Conn, _ = Config.DB.DB()

	Config.DB.AutoMigrate(&Models.Activity{}, &Models.Todo{})
	r := Routes.SetupRouter()

	//running
	r.Run(":3030")
}
