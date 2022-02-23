package models

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func GetDatabase() *gorm.DB {
	return DB
}

func ConnectDatabase() *gorm.DB {
	// var (
	// 	user     = os.Getenv("MYSQL_USERNAME") // e.g. 'my-db-user'
	// 	password = os.Getenv("MYSQL_PASSWORD") // e.g. 'my-db-password'
	// 	ip       = os.Getenv("MYSQL_IP")       // e.g. '127.0.0.1' ('172.17.0.1' if deployed to GAE Flex)
	// 	port     = os.Getenv("MYSQL_PORT")     // e.g. '3306'
	// 	database = os.Getenv("MYSQL_DATABASE") // e.g. 'my-database'
	// )
	// dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8mb4&parseTime=True&loc=Local", user, password, ip, port, database)
	dsn := "erica:c3dpdGFwcDA4MDEh@tcp(34.85.27.188:3306)/erica?parseTime=true&loc=Local"
	fmt.Println("1")
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	fmt.Println("2")
	db.AutoMigrate(&User{})
	// defer db.DB().Close()
	fmt.Println("3")
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(5 * time.Minute)
	DB = db
	return db
}
