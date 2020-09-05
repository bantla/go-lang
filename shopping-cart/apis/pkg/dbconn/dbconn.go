// Package dbconn (database connection) defines functions that access database
package dbconn

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ConnectMySQL function takes a data source name (dsn) to make a connection to RDMS
// (relational database management system) and
// returns a database connection (or an error if the connection is failed)
func ConnectMySQL(dsn string) (db *gorm.DB, err error){
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return
}
