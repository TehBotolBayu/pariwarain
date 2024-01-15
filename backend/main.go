package main

import (
	"github.com/TehBotolBayu/pariwarainAPI/routers"
	// "github.com/TehBotolBayu/pariwarainAPI/models"
	// "github.com/TehBotolBayu/pariwarainAPI/utils"
	// "fmt"
    // "database/sql"
    // "database/sql"
    // "fmt"
	// "os"
    _ "github.com/lib/pq"
	// "time"
	// "fmt"
	// "github.com/joho/godotenv"
	// "strconv"
	// "github.com/TehBotolBayu/pariwarainAPI/utils"
)

// func goDotEnvVariable(key string) string {
// 	// load .env file
// 	errorenv := godotenv.Load(".env")

// 	if errorenv != nil {
// 		panic("Error loading .env file")
// 	}

// 	return os.Getenv(key)
// }
 

func main() {
	routers.Router()
}
 
