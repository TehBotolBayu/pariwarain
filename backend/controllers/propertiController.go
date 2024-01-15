package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	// "mime/multipart"
	"net/http"
	// "time"

	"github.com/TehBotolBayu/pariwarainAPI/utils"
	_ "github.com/lib/pq"
	// "google.golang.org/protobuf/types/known/timestamppb"
)

type Properti struct {
	 Id string
	 Foto string
	 Judul string
	 Alamat string
	 Harga int64
	 Password string
	 Createdat string
	 Updatedat string
	 Pengguna_id string
}

func GetProperti(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query(`SELECT * FROM Properti`)
	utils.CheckError(err)
 
	defer rows.Close()
	for rows.Next() {
		var readProperti = new(Properti)
	
		err = rows.Scan(&readProperti.Id, &readProperti.Foto, 
			&readProperti.Judul, &readProperti.Alamat, &readProperti.Harga, 
			&readProperti.Password, &readProperti.Createdat, &readProperti.Updatedat, 
			&readProperti.Pengguna_id)

		utils.CheckError(err)
	
		fmt.Println(readProperti.Id, readProperti.Foto, readProperti.Judul, 
			readProperti.Alamat, readProperti.Harga, readProperti.Password, 
			readProperti.Createdat, readProperti.Updatedat, 
			readProperti.Pengguna_id)
	}
	
	utils.CheckError(err)
	
	fmt.Fprintf(w, "GET request received")
}

func PostProperti(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	var prop Properti
	
	json.Unmarshal(body, &prop)

	fmt.Println(prop.Id, prop.Foto, prop.Judul, prop.Alamat, prop.Harga, prop.Password, prop.Createdat, prop.Updatedat, prop.Pengguna_id)

	defer r.Body.Close()

	insertDynStmt := `INSERT INTO Properti( id, foto, judul, alamat, harga, password, createdat, updatedat, pengguna_id) 
	values($1, $2, $3, $4, $5, $6, $7, $8, $9)`
    _, e := db.Exec(insertDynStmt, prop.Id, prop.Foto, prop.Judul, prop.Alamat, prop.Harga, prop.Password, prop.Createdat, prop.Updatedat, prop.Pengguna_id)

	utils.CheckError(e)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "POST request received")
	
}

func PutProperti(w http.ResponseWriter, r *http.Request) {
	updateStmt := `update "Students" set "Name"=$1, "Roll"=$2 where "id"=$3`
	_, e := db.Exec(updateStmt, "Mary", 3, 2)
	fmt.Fprintf(w, "PUT request received")
	utils.CheckError(e)
}

func DeleteProperti(w http.ResponseWriter, r *http.Request) {
	deleteStmt := `delete from "Students" where id=$1`
	_, e := db.Exec(deleteStmt, 1)
	fmt.Fprintf(w, "DELETE request received")
	utils.CheckError(e)
}

