package models

import (
    "database/sql"
    "fmt"
	"os"
    _ "github.com/lib/pq"
	"github.com/joho/godotenv"
	"strconv"
	"github.com/TehBotolBayu/pariwarainAPI/utils"
)

func goDotEnvVariable(key string) string {
	// load .env file
	errorenv := godotenv.Load(".env")

	if errorenv != nil {
		panic("Error loading .env file")
	}

	return os.Getenv(key)
}
 

func Setup() *sql.DB {
	host     := goDotEnvVariable("host")
    port, _  := strconv.Atoi(goDotEnvVariable("port"))
    user     := goDotEnvVariable("user")
    password := goDotEnvVariable("password")
    dbname   := goDotEnvVariable("dbname")

        // connection string
    psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
         
        // open database
    db, err := sql.Open("postgres", psqlconn)
    utils.CheckError(err)
     
        // close database
    defer db.Close()
 
        // check db
    err = db.Ping()
    utils.CheckError(err)
  
    fmt.Println("Connected to Database!")
	return db
}