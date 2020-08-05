package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // postgres driver
)

const (
	// host is unexported
	host = "188.166.249.111"
	// port is unexported
	port = 5432
	// user is unexported
	user = "quochuy"
	// password is unexported
	password = "admin123"
	// dbname is unexported
	dbname = "adv_dev"
)

// ConnectDB opens a connection to the database
func ConnectDB() *sql.DB {
	// pg config
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err.Error())
	}
	pingError := db.Ping()
	if pingError != nil {
		panic(pingError.Error())
	}
	fmt.Println("Database connection established")
	return db
}
