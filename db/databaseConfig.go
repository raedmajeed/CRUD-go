package db

import (
	"fmt"
	dao "raedmajeed/entity"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


func ConnectDatabase() *gorm.DB {

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", 
		Username, Printlnassword, Host, Port, DBname,
		) 
						
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("Connection to DB %s Failed, Error: %s", DBname, err)
	}

	if err := database.AutoMigrate(dao.User{}); err != nil {
		fmt.Printf("Automigrate failed for %s failed", DBname)
	}

	return database
}