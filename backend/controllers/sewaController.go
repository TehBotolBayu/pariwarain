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

type Penyewaan struct {
	Id string
	PostId string
	PenyewaId string
	PemilikId string
	Start string
	End string
	Price int64
	Status string
	KodeUnik string
}

type DaftarSewa struct {
	Penyewaan []Penyewaan
}

func GetSewa(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query(`SELECT * FROM Penyewaan`)
	utils.CheckError(err)
	defer rows.Close()
 
	var sewaData DaftarSewa 

	for rows.Next() {
		var readSewa = new(Penyewaan)
		
		err = rows.Scan(&readSewa.Id, &readSewa.PostId, 
			&readSewa.PenyewaId, &readSewa.PemilikId, &readSewa.Start, 
			&readSewa.End, &readSewa.Price, &readSewa.Status, &readSewa.KodeUnik)

		utils.CheckError(err)

		sewaData.Penyewaan = append(sewaData.Penyewaan, *readSewa)
	}

	jsonResponse, err := json.Marshal(sewaData)
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

func PostSewa(w http.ResponseWriter, r *http.Request) {
	var sewa Penyewaan
	
	// json.Unmarshal(body, &prop)
	err := r.ParseMultipartForm(10 << 20) // 10 MB limit for file size
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		panic(err)
	}
	r.ParseForm()
	
	// Id = r.FormValue("")
	sewa.PostId = r.FormValue("postId")
	sewa.PenyewaId = r.FormValue("penyewaId")
	sewa.PemilikId = r.FormValue("PemilikId")
	sewa.Start = r.FormValue("start")
	sewa.End = r.FormValue("end")
	sewa.Price, _ = strconv.ParseInt(r.FormValue("price"), 10, 64)
	sewa.Status = r.FormValue("status")
	sewa.KodeUnik = r.FormValue("KodeUnik")

	// fmt.Println(prop.Id, prop.Foto, prop.Judul, prop.Alamat, prop.Harga, prop.Createdat, prop.Updatedat, prop.Pengguna_id)

	defer r.Body.Close()

	insertDynStmt := `INSERT INTO Penyewaan(postid, penyewaid, pemilikid, startrent, endrent, price, status, kodeunik) 
	values($1, $2, $3, $4, $5, $6, $7, $8)`
    _, e := db.Exec(insertDynStmt, sewa.PostId, sewa.PenyewaId, sewa.PemilikId, sewa.Start, sewa.End, sewa.Price, sewa.Status, sewa.KodeUnik)

	utils.CheckError(e)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "POST request received")
}

func PutSewa(w http.ResponseWriter, r *http.Request) {
	var sewa Penyewaan
	
	// json.Unmarshal(body, &prop)
	err := r.ParseMultipartForm(10 << 20) // 10 MB limit for file size
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		panic(err)
	}
	r.ParseForm()
	
	// Get a reference to the file from the request
	sewa.PostId = r.FormValue("postId")
	sewa.PenyewaId = r.FormValue("penyewaId")
	sewa.PemilikId = r.FormValue("PemilikId")
	sewa.Start = r.FormValue("start")
	sewa.End = r.FormValue("end")
	sewa.Price, _ = strconv.ParseInt(r.FormValue("price"), 10, 64)
	sewa.Status = r.FormValue("status")
	sewa.KodeUnik = r.FormValue("KodeUnik")

	var e error
	updateStmt := `update Penyewaan postid = $1, penyewaid = $2, 
	pemilikid = $3, startrent = $4, endrent = $5, price = $6, status = $7, kodeunik = $8
	where id=$9`
	_, e = db.Exec(updateStmt, sewa.PostId, sewa.PenyewaId, sewa.PemilikId, sewa.Start, sewa.End, sewa.Price, sewa.Status, sewa.KodeUnik, sewa.Id)
	
	fmt.Fprintf(w, "PUT request received")
	utils.CheckError(e)  
}

func DeleteSewa(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	id := queryParams.Get("id")

	deleteStmt := `DELETE FROM Penyewaan WHERE id=$1`
	_, e := db.Exec(deleteStmt, id)
	fmt.Fprintf(w, "DELETE request received")
	utils.CheckError(e)
}