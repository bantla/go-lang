package main

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
  gorm.Model
  Name         string
  Age          sql.NullInt64
  Birthday     *time.Time
  Email        string  `gorm:"type:varchar(100);unique_index"`
  Role         string  `gorm:"size:255"` // set field size to 255
  MemberNumber *string `gorm:"unique;not null"` // set member number to unique and not null
  Num          int     `gorm:"AUTO_INCREMENT"` // set num to auto incrementable
  Address      string  `gorm:"index:addr"` // create index with name `addr` for address
  IgnoreMe     int     `gorm:"-"` // ignore this field
}

func main() {
  db, err := gorm.Open("mysql", "root:root-admin@(localhost)/gorm_demo?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()

	if err != nil {
		fmt.Println(err)
		return
	}

	birthday := time.Now()

	db.AutoMigrate(&User{})

	memberNumber := "Hi"

	// http://gorm.io/docs/migration.html
	// Create table for model `User`
	// db.CreateTable(&User{})

	user := User{Name: "BIN", Age: sql.NullInt64{0, false}, Birthday: &birthday, MemberNumber: &memberNumber, Email: "h@h"}

	var users []User
	var ages []int64

	fmt.Println(user)
	db.Find(&users).Pluck("age", &ages)
	fmt.Println(users)
	fmt.Println(ages)
	// fmt.Println(db.NewRecord(user))
	// fmt.Println(db.Create(&user))
	// fmt.Println(db.NewRecord(user))
}
