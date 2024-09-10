package infrastructure

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

type Config struct {
	host     string
	port     string
	user     string
	password string
	dbname   string
}

var c = Config{
	host:     "",
	port:     "5432",
	user:     "",
	password: "",
	dbname:   "",
}

func Connect() *sql.DB {
	port, err := strconv.Atoi(c.port)
	connInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		c.host, port, c.user, c.password, c.dbname)

	db, err := sql.Open("postgres", connInfo)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected to data base!")
	return db
}


func ParamConnection() {
	c.host = os.Getenv("DATABASE_HOST")
	c.password = os.Getenv("DATABASE_PASS")
	c.user = os.Getenv("DATABASE_USER")
	c.dbname = os.Getenv("DATABASE_NAME")
	if os.Getenv("DATABASE_PORT") != "" {
		c.port = os.Getenv("DATABASE_PORT")
	}

	if c.host == "" {
		log.Fatal("no host for the dataBase")
	}
	if c.password == "" {
		log.Fatal("no password for the dataBaseb")
	}
	if c.user == "" {
		log.Fatal("no user for the dataBase")
	}
	if c.dbname == "" {
		log.Fatal("no name for the dataBase")
	}
}

