package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	_ "github.com/go-sql-driver/mysql"
)

type EmailDetails struct {
	Id      int
	Address string
	Subject string
	Message string
}

// Oh, would you like a database? Here's how you can fire it up!
// create database form_persistance;
// create user 'go-squee' identified by 'my-new-password';
// GRANT ALL PRIVILEGES on *.* to 'go-squee';
// mysql> create table emails (id INT NOT NULL AUTO_INCREMENT, address VARCHAR(512), subject VARCHAR(1024), message TEXT, PRIMARY KEY ( id ));

var DB *sql.DB

func main() {
	sqlConnect()
	defer DB.Close()
	r := mux.NewRouter()

	r.HandleFunc("/", handleRoot)
	r.HandleFunc("/emails", handleEmails)
	r.HandleFunc("/emails/{emailId}", handleEmail)

	log.Fatal(http.ListenAndServe(":8080", r))
}

func sqlConnect() {
	db, err := sql.Open("mysql", "go-squee:my-new-password@(127.0.0.1:3306)/form_persistance?parseTime=true")

	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	DB = db
}

type orderData struct {
	Name        string
	OrderNumber int
	ShipDate    string
}

func newCrap(w http.ResponseWriter, r *http.Request) {
	var header = `
		WidgetCo, Ltd.
		463 Shoe Factory Rd.
		Hamford, VT 20202
	`
	var footer = `
		Thank you for your business,
		WidgetCo Order Fulfillment Department
		Ph: 818-555-0123 Email: orders@widgetco.com
	`
	var thanks = `
	{{ template "header" }}
	<h1>Contact</h1>
	<form method="POST">
		<label>Address:</label><br />
		<input type="text" name="address" /><br />
		<label>Subject:</label><br />
		<input type="text" name="subject" /><br />
		<label>Message:</label><br />
		<textarea name="message"></textarea><br />
		<input type="submit" />
	</form>
	{{ template "footer"  }}
	`

	t, _ := template.New("header").Parse(header)
	t.New("footer").Parse(footer)
	t.New("thanks").Parse(thanks)
	t.ExecuteTemplate(w, "thanks", orderData{"Sleve McDichael", 17104, "2018-10-10"})
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	t, _ := template.New("form").ParseFiles("templates/forms.html", "templates/footer.html")
	// t.New("footer").ParseFiles("templates/footer.html")
	// tmpl := template.Must(template.ParseFiles("templates/forms.html"))
	if r.Method == http.MethodGet {
		t.ExecuteTemplate(w, "form", nil)
		return
	}
	// tmpl.New("footer").ParseFiles("templates/footer.html")

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
		t.ExecuteTemplate(w, "form", struct{ Success bool }{true})
		return
	}

	response := "lmao that method is totes not supported. good luck 2 ya!"
	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write([]byte(response))
	fmt.Fprintf(w, response)
	log.Println(response)

}

func handleEmails(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/emails.html"))

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

	// query := "SELECT id, username, password, created_at FROM users WHERE id = ?"
	// if err := db.QueryRow(query, 2).Scan(&id, &username, &password, &createdAt); err != nil {
	email := EmailDetails{}
	query := "SELECT id, address, subject, message FROM emails WHERE id = ?"
	if err := DB.QueryRow(query, emailId).Scan(&email.Id, &email.Address, &email.Subject, &email.Message); err != nil {
		tmpl := template.Must(template.ParseFiles("templates/404.html"))
		tmpl.Execute(w, nil)
	} else {
		tmpl := template.Must(template.ParseFiles("templates/email.html"))
		tmpl.Execute(w, email)
	}
}
