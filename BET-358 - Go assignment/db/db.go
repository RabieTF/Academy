package DB

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

var Connection *sql.DB

func ConnectToDB() {
	cfg := mysql.Config{
		User:                 os.Getenv("DBUSER"),
		Passwd:               os.Getenv("DBPASS"),
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "Shopping",
		AllowNativePasswords: true,
	}

	var err error

	Connection, err = sql.Open("mysql", cfg.FormatDSN())

	if err != nil {
		log.Fatal(err)
	}

	pingErr := Connection.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
}
