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

type Daftar struct {
	Pengguna []Pengguna
}

func Login(w http.ResponseWriter, r *http.Request) {
	//get data from request
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}
	fmt.Println("Request Body:", string(body))

	var user Pengguna

	json.Unmarshal(body, &user)

	defer r.Body.Close()

	// check if email exist
	rows, err := db.Query(`SELECT * FROM Pengguna WHERE email=$1`, user.Email)
	utils.CheckError(err)
	defer rows.Close()

	// var daftarPengguna Daftar
 
	var id string
	var nama string
	var email string
	var pw string

	for rows.Next() {
		err = rows.Scan(&id, &nama, &email, &pw)
		utils.CheckError(err)
	}

	
	// check if password is correct
	if(pw == ""){
		fmt.Fprintf(w, "Account is not registered")
	} else if(user.Pw == pw){
		fmt.Fprintf(w, "Welcome "+nama)
	} else {
		fmt.Fprintf(w, "Wrong Password")
	}
	
}

func GetPengguna(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query(`SELECT * FROM Pengguna`)
	utils.CheckError(err)
	defer rows.Close()

	var daftarPengguna Daftar
 
	for rows.Next() {
		var name string
		var name2 string
		var name3 string
		var name4 string
		
		err = rows.Scan(&name, &name2, &name3, &name4)
		utils.CheckError(err)
		
		person := Pengguna{
			Id: name,
			Nama: name2,
			Email: name3,
			Pw: name4,
		}
		daftarPengguna.Pengguna = append(daftarPengguna.Pengguna, person)

		// fmt.Println(name, name2, name3, name4)
	}

	jsonResponse, err := json.Marshal(daftarPengguna)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	// Set the Content-Type header to indicate JSON content
	w.Header().Set("Content-Type", "application/json")
	// Write the JSON response to the client
	w.Write(jsonResponse)
	
	utils.CheckError(err)
	
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

	insertDynStmt := `INSERT INTO Pengguna(nama, email, pw) values($1, $2, $3)`
    _, e := db.Exec(insertDynStmt, user.Email, user.Nama, user.Pw)
	fmt.Fprintf(w, "User register success")
	utils.CheckError(e)
}

func PutPengguna(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}
	fmt.Println("Request Body:", string(body))

	var user Pengguna

	json.Unmarshal(body, &user)

	defer r.Body.Close()

	updateStmt := `update Pengguna set nama=$1, email=$2, pw=$3 where id=$4`
	_, e := db.Exec(updateStmt, user.Nama, user.Email, user.Pw, user.Id)
	fmt.Fprintf(w, "PUT request received")
	utils.CheckError(e)
}

func DeletePengguna(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	id := queryParams.Get("id")

	deleteStmt := `delete from Pengguna where id=$1`
	_, e := db.Exec(deleteStmt, id)
	fmt.Fprintf(w, "DELETE request received")
	utils.CheckError(e)
}