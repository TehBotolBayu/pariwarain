package controllers

// import (
// 	"net/http"
// 	"fmt"
// 	_ "github.com/lib/pq"		
// )

// func handleGet(w http.ResponseWriter, r *http.Request) {
// 	rows, err := db.Query(`SELECT "Name", "Roll" FROM "Students"`)
// 	CheckError(err)
	
// 	defer rows.Close()
// 	for rows.Next() {
// 		var name string
// 		var roll int
	
// 		err = rows.Scan(&name, &roll)
// 		CheckError(err)
	
// 		fmt.Println(name, roll)
// 	}
// 	fmt.Fprintf(w, "GET request received")
// }

// func handlePost(w http.ResponseWriter, r *http.Request) {
// 	insertDynStmt := `insert into "Students"("Name", "Roll") values($1, $2)`
//     _, e = db.Exec(insertDynStmt, "Jane", 2)
// 	fmt.Fprintf(w, "POST request received")
// }

// func handlePut(w http.ResponseWriter, r *http.Request) {
// 	updateStmt := `update "Students" set "Name"=$1, "Roll"=$2 where "id"=$3`
// 	_, e := db.Exec(updateStmt, "Mary", 3, 2)
// 	fmt.Fprintf(w, "PUT request received")
// }

// func handleDelete(w http.ResponseWriter, r *http.Request) {
// 	deleteStmt := `delete from "Students" where id=$1`
// 	_, e := db.Exec(deleteStmt, 1)
// 	fmt.Fprintf(w, "DELETE request received")
// }
