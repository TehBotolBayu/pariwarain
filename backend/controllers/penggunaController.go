package controllers

import (
	"net/http"
    "database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/TehBotolBayu/pariwarainAPI/utils"
	"github.com/TehBotolBayu/pariwarainAPI/models"
	"io"
)

var db *sql.DB = models.Setup()

func HandleGet(w http.ResponseWriter, r *http.Request) {
	// x := r.Body.Read("asd ");

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}
	fmt.Println("Request Body:", string(body))

	defer r.Body.Close()


	rows, errorDB := db.Query(`SELECT "Name", "Roll" FROM "Students"`)
	utils.CheckError(errorDB)
	 
	defer rows.Close()
	for rows.Next() {
		var id string
		var name string
		var email string
		var pw string
	
		err = rows.Scan(&id, &name, &email, &pw)
		utils.CheckError(err)
	
		fmt.Println(id, name, email, pw)
	}
	fmt.Fprintf(w, "GET request received")
}

func HandlePost(w http.ResponseWriter, r *http.Request) {
	insertDynStmt := `insert into "Students"("Name", "Roll") values($1, $2)`
    _, e := db.Exec(insertDynStmt, "Jane", 2)
	fmt.Fprintf(w, "POST request received")
	utils.CheckError(e)
}

func HandlePut(w http.ResponseWriter, r *http.Request) {
	updateStmt := `update "Students" set "Name"=$1, "Roll"=$2 where "id"=$3`
	_, e := db.Exec(updateStmt, "Mary", 3, 2)
	fmt.Fprintf(w, "PUT request received")
	utils.CheckError(e)

}

func HandleDelete(w http.ResponseWriter, r *http.Request) {
	deleteStmt := `delete from "Students" where id=$1`
	_, e := db.Exec(deleteStmt, 1)
	fmt.Fprintf(w, "DELETE request received")
	utils.CheckError(e)

}
