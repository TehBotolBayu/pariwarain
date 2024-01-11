package routers

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// )

// func route() {
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		switch r.Method {
// 		case http.MethodGet:
// 			handleGet(w, r)
// 		case http.MethodPost:
// 			handlePost(w, r)
// 		case http.MethodPut:
// 			handlePut(w, r)
// 		case http.MethodDelete:
// 			handleDelete(w, r)
// 		default:
// 			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
// 		}
// 	})

// 	fmt.Println("Server is running on http://localhost:8080")
// 	log.Fatal(http.ListenAndServe(":8080", nil))
// }
