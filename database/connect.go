package database

import (
	"os"

	//"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/lib/pq"
)

var DB *gorm.DB
type User struct {
	Id    int    `json:"id" param:"id"`
	Name  string `json:"name"`
}
func Connect() {
	//err := godotenv.Load()
	//if err != nil {
	//	panic(err.Error())
	//}

	url := os.Getenv("DATABASE_URL")
	dsn, err := pq.ParseURL(url)
	if err != nil {
		panic(err.Error())
	}
	dsn += " sslmode=require"

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	DB.AutoMigrate(&User{})
}
