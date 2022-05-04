package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	connector "example.com/sql-utils"
	"github.com/gorilla/mux"

	_ "github.com/go-sql-driver/mysql"
)

type EmailDetails struct {
	Id      int
	Address string
	Subject string
	Message string
}

var DB *sql.DB

func main() {
	DB = connector.Connect()
	defer DB.Close()
	r := mux.NewRouter()

	r.HandleFunc("/", handleRoot)
	r.HandleFunc("/emails", handleEmails)
	r.HandleFunc("/emails/{emailId}", handleEmail)

	log.Fatal(http.ListenAndServe(":8080", r))
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("forms.html"))
	if r.Method == http.MethodGet {
		tmpl.Execute(w, nil)
		return
	}

	if r.Method == http.MethodPost {
		details := EmailDetails{
			Address: r.FormValue("address"),
			Subject: r.FormValue("subject"),
			Message: r.FormValue("message"),
		}
		_, err := DB.Exec("INSERT INTO emails (address, subject, message) VALUES (?, ?, ?)", details.Address, details.Subject, details.Message)
		if err != nil {
			log.Fatal(err)
		}

		tmpl.Execute(w, struct{ Success bool }{true})
		return
	}

	response := "lmao that method is totes not supported. good luck 2 ya!"
	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write([]byte(response))
	fmt.Fprintf(w, response)
	log.Println(response)
}

func handleEmails(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("emails.html"))

	query := "SELECT * FROM emails;"
	rows, err := DB.Query(query)
	if err != nil {
		log.Fatal(err)
	}

	details := []EmailDetails{}

	for rows.Next() {
		email := EmailDetails{}
		err := rows.Scan(&email.Id, &email.Address, &email.Subject, &email.Message)
		if err != nil {
			log.Fatal(err)
		}
		details = append(details, email)
	}

	tmpl.Execute(w, details)
}

func handleEmail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	emailId := vars["emailId"]

	email := EmailDetails{}
	query := "SELECT id, address, subject, message FROM emails WHERE id = ?"
	if err := DB.QueryRow(query, emailId).Scan(&email.Id, &email.Address, &email.Subject, &email.Message); err != nil {
		tmpl := template.Must(template.ParseFiles("404.html"))
		tmpl.Execute(w, nil)
	} else {
		tmpl := template.Must(template.ParseFiles("email.html"))
		tmpl.Execute(w, email)
	}
}
