package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"github.com/TehBotolBayu/pariwarainAPI/models"
	"github.com/TehBotolBayu/pariwarainAPI/utils"
	_ "github.com/lib/pq"
)

var db *sql.DB = models.Setup()

type Pengguna struct {
	Id string
	Nama string
	Email string
	Pw string
}

func GetPengguna(w http.ResponseWriter, r *http.Request) {
	// x := r.Body.Read("asd ");

	rows, err := db.Query(`SELECT * FROM Pengguna`)
	utils.CheckError(err)
 
	defer rows.Close()
	for rows.Next() {
		var name string
		var name2 string
		var name3 string
		var name4 string
		// var roll int
	
		err = rows.Scan(&name, &name2, &name3, &name4)
		utils.CheckError(err)
	
		fmt.Println(name, name2, name3, name4)
	}
	
	utils.CheckError(err)
	
	fmt.Fprintf(w, "GET request received")
}

func PostPengguna(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}
	fmt.Println("Request Body:", string(body))

	var user Pengguna

	json.Unmarshal(body, &user)

	defer r.Body.Close()

	insertDynStmt := `INSERT INTO Pengguna(id, nama, email, pw) values($1, $2, $3, $4)`
    _, e := db.Exec(insertDynStmt, user.Id, user.Email, user.Nama, user.Pw)
	fmt.Fprintf(w, "User register success")
	utils.CheckError(e)
}

func PutPengguna(w http.ResponseWriter, r *http.Request) {

	// updateStmt := `update "Students" set "Name"=$1, "Roll"=$2 where "id"=$3`
	// _, e := db.Exec(updateStmt, "Mary", 3, 2)
	fmt.Fprintf(w, "PUT request received")
	// utils.CheckError(e)
}

func DeletePengguna(w http.ResponseWriter, r *http.Request) {
	// deleteStmt := `delete from "Students" where id=$1`
	// _, e := db.Exec(deleteStmt, 1)
	fmt.Fprintf(w, "DELETE request received")
	// utils.CheckError(e)
}