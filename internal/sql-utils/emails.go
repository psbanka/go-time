package sqlutils

import (
	"database/sql"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type EmailDetails struct {
	Id      int
	Address string
	Subject string
	Message string
}

func CreateEmail(r *http.Request, DB *sql.DB) (EmailDetails, error) {
	details := EmailDetails{
		Address: r.FormValue("address"),
		Subject: r.FormValue("subject"),
		Message: r.FormValue("message"),
	}
	_, err := DB.Exec("INSERT INTO emails (address, subject, message) VALUES (?, ?, ?)", details.Address, details.Subject, details.Message)
	return details, err
}
