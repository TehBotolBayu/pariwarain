package routers

import (
	"fmt"
	"log"
	"net/http"
	"github.com/TehBotolBayu/pariwarainAPI/controllers"
)

func PenggunaRouter() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			controllers.HandleGet(w, r)
		// case http.MethodPost:
		// 	controllers.HandlePost(w, r)
		case http.MethodPut:
			controllers.HandlePut(w, r)
		case http.MethodDelete:
			controllers.HandleDelete(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

