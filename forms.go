package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type ContactDetails struct {
	Email   string
	Subject string
	Message string
}

func main() {
	http.HandleFunc("/", handleRoot)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("forms.html"))

  if r.Method == http.MethodGet {
    tmpl.Execute(w, nil)
    return
  }

  if r.Method == http.MethodPost {

    details := ContactDetails{
      Email:   r.FormValue("email"),
      Subject: r.FormValue("subject"),
      Message: r.FormValue("message"),
    }

    // do something with details
    _ = details

    // Things to do to actually persist this to DB
    // 1. Create the database on the local machine
    // 2. Create a user for the database
    // 3. Setup appropriate tables (in our case, just one for now)
    // 4. Connect to the database in this method
    // 5. Actually persist data
    // 6. [OPTIONAL] Add a new route with HTTP BASIC AUTH to get access to the
    // data.
    // 7. ...
    // 8. Profit!

    tmpl.Execute(w, struct{ Success bool }{true})
    return
  }

  response := "lmao that method is totes not supported. good luck 2 ya!"
  w.WriteHeader(http.StatusMethodNotAllowed)
  w.Write([]byte(response))
  fmt.Fprintf(w, response)
  log.Println(response)

}
