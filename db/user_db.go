package db

import (
	"database/sql"
	
	"fmt"
	"os"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type User struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
}
var (
	Client *sql.DB
	username = "root"
	hostname = "localhost:3306"
	databasename = "go_test"
)


func init() {
	if envErr := godotenv.Load(); envErr != nil {
		panic("not get env")
	}
	password := os.Getenv("DB_PASSWORD")
	dbIndo := fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, databasename)
	
  var err error
	Client, err = sql.Open("mysql", dbIndo)
	if err != nil {
		panic(err)
	}
	if clientErr := Client.Ping(); clientErr != nil {
		panic(clientErr)
	}
	fmt.Println("db on")
}