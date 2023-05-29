package databases

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Conn() {
	// dsn := os.Getenv("DB")
	database, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/go_restapi_gin?parseTime=true"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB = database

}
