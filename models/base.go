package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB //база данных

func init() {
	conn, err := gorm.Open("postgres", "user=admin_shop password=123456 dbname=shop sslmode=disable")
	if err != nil {
		fmt.Print(err)
	}

	db = conn
	db.Debug().AutoMigrate(&Users{}, &Product{})

	initManagers()
}

func initManagers() {
	managers := make([]Users, 2)
	managers = append(managers, Users{Account: "manager1", Password: "qwerty123", Type: Manager})
	managers = append(managers, Users{Account: "manager2", Password: "qwerty321", Type: Manager})

	for _, val := range managers {
		account := val
		_ = account.Create()

	}
}

func GetDB() *gorm.DB {
	return db
}
