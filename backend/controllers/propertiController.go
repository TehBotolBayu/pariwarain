package controllers

import (
	"encoding/json"
	"fmt"
	// "io"
	"strconv"
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
	 Createdat string
	 Updatedat string
	 Pengguna_id string
}

type DaftarProperti struct {
	Properti []Properti
}

func GetProperti(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query(`SELECT * FROM Properti`)
	utils.CheckError(err)
	defer rows.Close()
 
	var propertiData DaftarProperti 

	for rows.Next() {
		var readProperti = new(Properti)
	
		err = rows.Scan(&readProperti.Id, &readProperti.Foto, 
			&readProperti.Judul, &readProperti.Alamat, &readProperti.Harga, 
			&readProperti.Createdat, &readProperti.Updatedat, 
			&readProperti.Pengguna_id)

		utils.CheckError(err)

		propertiData.Properti = append(propertiData.Properti, *readProperti)
	
		// fmt.Println(readProperti.Id, readProperti.Foto, readProperti.Judul, 
		// 	readProperti.Alamat, readProperti.Harga, readProperti.Password, 
		// 	readProperti.Createdat, readProperti.Updatedat, 
		// 	readProperti.Pengguna_id)

	}

	jsonResponse, err := json.Marshal(propertiData)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	// Set the Content-Type header to indicate JSON content
	w.Header().Set("Content-Type", "application/json")
	// Write the JSON response to the client
	w.Write(jsonResponse)
	
	utils.CheckError(err)
	
	// fmt.Fprintf(w, "GET request received")
}

func PostProperti(w http.ResponseWriter, r *http.Request, filename string) {
	var prop Properti
	
	// json.Unmarshal(body, &prop)
	err := r.ParseMultipartForm(10 << 20) // 10 MB limit for file size
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		panic(err)
	}
	r.ParseForm()
	
	// Get a reference to the file from the request
	prop.Id = r.FormValue("Id")
	prop.Foto = filename
	prop.Judul = r.FormValue("Judul")
	prop.Alamat = r.FormValue("Alamat")
	prop.Harga, _ = strconv.ParseInt(r.FormValue("Harga"), 10, 64)
	prop.Createdat = r.FormValue("Createdat")
	prop.Updatedat = r.FormValue("Updatedat")
	prop.Pengguna_id = r.FormValue("Pengguna_id")

	fmt.Println(prop.Id, prop.Foto, prop.Judul, prop.Alamat, prop.Harga, prop.Createdat, prop.Updatedat, prop.Pengguna_id)

	defer r.Body.Close()

	insertDynStmt := `INSERT INTO Properti(foto, judul, alamat, harga, createdat, updatedat, penggunaId) 
	values($1, $2, $3, $4, $5, $6, $7)`
    _, e := db.Exec(insertDynStmt, prop.Foto, prop.Judul, prop.Alamat, prop.Harga, prop.Createdat, prop.Updatedat, prop.Pengguna_id)

	utils.CheckError(e)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "POST request received")
}

func PutProperti(w http.ResponseWriter, r *http.Request, filename string) {
	var prop Properti
	
	// json.Unmarshal(body, &prop)
	err := r.ParseMultipartForm(10 << 20) // 10 MB limit for file size
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		panic(err)
	}
	r.ParseForm()
	
	// Get a reference to the file from the request
	prop.Id = r.FormValue("Id")
	prop.Foto = filename
	prop.Judul = r.FormValue("Judul")
	prop.Alamat = r.FormValue("Alamat")
	prop.Harga, _ = strconv.ParseInt(r.FormValue("Harga"), 10, 64)
	prop.Createdat = r.FormValue("Createdat")
	prop.Updatedat = r.FormValue("Updatedat")
	prop.Pengguna_id = r.FormValue("Pengguna_id")

	var e error
	updateStmt := `update Properti set foto=$1, judul=$2, alamat=$3, harga=$4, password=$5, updatedat=$6 where id=$7`
	if filename == "" {
		updateStmt = `update Properti set judul=$1, alamat=$2, harga=$3, updatedat=$4 where id=$5`
		_, e = db.Exec(updateStmt, prop.Judul, prop.Alamat, prop.Harga, prop.Updatedat, prop.Id)	
	} else {
		_, e = db.Exec(updateStmt, prop.Foto, prop.Judul, prop.Alamat, prop.Harga, prop.Updatedat, prop.Id)
	}
	fmt.Fprintf(w, "PUT request received")
	utils.CheckError(e)  
}

func DeleteProperti(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	id := queryParams.Get("id")

	deleteStmt := `DELETE FROM Properti WHERE id=$1`
	_, e := db.Exec(deleteStmt, id)
	fmt.Fprintf(w, "DELETE request received")
	utils.CheckError(e)
}