package models

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	// dsn := "root:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db_user := "root"
	db_password := ""
	db_name := "nest"
	db_port := "3306"
	db_host := ""

	dsn := db_user + ":" + db_password + "@tcp(" + db_host + ":" + db_port + ")/" + db_name + "?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Printf("%s%", err)
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&Post{}) // register Post model

	DB = database
}
