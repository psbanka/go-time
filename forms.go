package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type ContactDetails struct {
	Address string
	Subject string
	Message string
}

func main() {
	http.HandleFunc("/", handleRoot)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("forms.html"))
	db, err := sql.Open("mysql", "go-squee:my-new-password@(127.0.0.1:3306)/form_persistance?parseTime=true")

	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	if r.Method == http.MethodGet {
		tmpl.Execute(w, nil)
		return
	}

	if r.Method == http.MethodPost {

		details := ContactDetails{
			Address: r.FormValue("address"),
			Subject: r.FormValue("subject"),
			Message: r.FormValue("message"),
		}

		_, err := db.Exec("INSERT INTO emails (address, subject, message) VALUES (?, ?, ?)", details.Address, details.Subject, details.Message)
		if err != nil {
			log.Fatal(err)
		}

		// Things to do to actually persist this to DB
		// 1. Create the database on the local machine - DONE
		// 2. Create a user for the database - DONE
		// 3. Setup appropriate tables (in our case, just one for now) - DONE
		// 4. Connect to the database in this method - DONE
		// 5. Actually persist data - DONE
		// 5b. read all emails?
		// 5c. read an email?
		// 6. [OPTIONAL] Add a new route with HTTP BASIC AUTH
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
