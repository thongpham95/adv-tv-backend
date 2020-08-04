package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // postgres driver
)

const (
	// host is exported
	host = "localhost"
	// port is exported
	port = 5432
	// user is exported
	user = "quochuy"
	// password is exported
	password = "admin123"
	// dbname is exported
	dbname = "postgres"
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
