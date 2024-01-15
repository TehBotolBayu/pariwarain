package routers

import (
	"fmt"
	"log"
	"net/http"
	"github.com/TehBotolBayu/pariwarainAPI/controllers"
)

func Router() {
	http.HandleFunc("/pengguna", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			controllers.GetPengguna(w, r)
		case http.MethodPost:
			controllers.PostPengguna(w, r)
		case http.MethodPut:
			controllers.PutPengguna(w, r)
		case http.MethodDelete:
			controllers.DeletePengguna(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/properti", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			controllers.GetProperti(w, r)
		case http.MethodPost:
			var filename string = controllers.UploadImage(w, r)
			controllers.PostProperti(w, r, filename)
		case http.MethodPut:
			var fileUpdate string = controllers.UploadImage(w, r)
			controllers.PutProperti(w, r, fileUpdate)
		case http.MethodDelete:
			controllers.DeleteProperti(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			controllers.UploadImage(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}