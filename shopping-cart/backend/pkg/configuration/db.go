package configuration

import "fmt"

// Database defines properties used for accessing database
type Database struct {
	// Username is used to log in database
	Username string

	// Password is used to log in database
	Password string

	// DBName is the name of database working on
	DBName string

	// Protocol is the protocol of database
	Protocol string

	// Host is the database host
	Host string

	// Port is the database port
	Port int
}

// GetConnectionURI method returns the connection uri
func (db Database) GetConnectionURI() string {
	uri := fmt.Sprintf(
		"%v:%v@%v(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		db.Username,
		db.Password,
		db.Protocol,
		db.Host,
		db.Port,
		db.DBName,
	)

	return uri
}
