package pq

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"time"
)

var db *gorm.DB

func conn() (dbErr error) {
	dsn := os.Getenv("PG_URL")
	db, dbErr = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return
}

func GetConn() *gorm.DB {
	for  {
		if db != nil {
			return db
		}
		fmt.Println("fail to get db, connect db again")
		if err := conn(); err != nil {
			fmt.Printf("conn fail: %v\n", err.Error())
			time.Sleep(3*time.Second)
		}
	}
}
