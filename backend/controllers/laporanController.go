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

type Laporan struct {
	Id string 
	PenyewaanId string 
	Foto string 
	Createdat string 
	Deskripsi string
}

type DaftarLaporan struct {
	Laporan []Laporan
}

func GetLaporan(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	sewaId := queryParams.Get("sewaId")

	rows, err := db.Query(`SELECT * FROM Laporan WHERE penyewaanid=$1`, sewaId)
	utils.CheckError(err)
	defer rows.Close()
 
	var laporanData DaftarLaporan

	for rows.Next() {
		var readLaporan = new(Laporan)
	
		err = rows.Scan(&readLaporan.Id, &readLaporan.PenyewaanId, 
			&readLaporan.Foto, &readLaporan.Createdat, &readLaporan.Deskripsi)

		utils.CheckError(err)

		laporanData.Laporan = append(laporanData.Laporan, *readLaporan)
	
		// fmt.Println(readProperti.Id, readProperti.Foto, readProperti.Judul, 
		// 	readProperti.Alamat, readProperti.Harga, readProperti.Password, 
		// 	readProperti.Createdat, readProperti.Updatedat, 
		// 	readProperti.Pengguna_id)

	}

	jsonResponse, err := json.Marshal(laporanData)
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

func PostLaporan(w http.ResponseWriter, r *http.Request, filename string) {
	var lapor Laporan
	
	// json.Unmarshal(body, &prop)
	err := r.ParseMultipartForm(10 << 20) // 10 MB limit for file size
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		panic(err)
	}
	r.ParseForm()
	
	// Get a reference to the file from the request
	lapor.Foto = filename
	lapor.Deskripsi = r.FormValue("Deskripsi")
	lapor.PenyewaanId = r.FormValue("PenyewaanId")
	lapor.Createdat = r.FormValue("CreatedAt")

	// fmt.Println(prop.Id, prop.Foto, prop.Judul, prop.Alamat, prop.Harga, prop.Createdat, prop.Updatedat, prop.Pengguna_id)

	defer r.Body.Close()

	insertDynStmt := `INSERT INTO Laporan(penyewaanid, foto, createdat, deskripsi) 
	values($1, $2, $3, $4)`

    _, e := db.Exec(insertDynStmt, lapor.PenyewaanId, lapor.Foto, lapor.Createdat, lapor.Deskripsi)

	utils.CheckError(e)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "POST request received")
}

func PutLaporan(w http.ResponseWriter, r *http.Request, filename string) {
	var lapor Laporan
	
	// json.Unmarshal(body, &prop)
	err := r.ParseMultipartForm(10 << 20) // 10 MB limit for file size
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		panic(err)
	}
	r.ParseForm()
	
	// Get a reference to the file from the request
	lapor.Foto = filename
	lapor.Deskripsi = r.FormValue("Deskripsi")
	lapor.PenyewaanId = r.FormValue("PenyewaanId")
	lapor.Createdat = r.FormValue("CreatedAt")

	var e error
	updateStmt := `update Properti set penyewaanid=$1, foto=$2, createdat=$3, deskripsi=$4 where id=$5`
	if filename == "" {
		updateStmt = `update Properti set penyewaanid=$1, createdat=$2, deskripsi=$3 where id=$4`
		_, e = db.Exec(updateStmt, lapor.PenyewaanId, lapor.Createdat, lapor.Deskripsi, lapor.Id)	
	} else {
		_, e = db.Exec(updateStmt, lapor.PenyewaanId, lapor.Foto, lapor.Createdat, lapor.Deskripsi, lapor.Id)
	}
	fmt.Fprintf(w, "PUT request received")
	utils.CheckError(e)  
}

func DeleteLaporan(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	id := queryParams.Get("id")

	deleteStmt := `DELETE FROM Laporan WHERE id=$1`
	_, e := db.Exec(deleteStmt, id)
	fmt.Fprintf(w, "DELETE request received")
	utils.CheckError(e)
}