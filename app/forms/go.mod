module go-time/forms

go 1.16

require (
	example.com/sql-utils v0.0.0-00010101000000-000000000000 // indirect
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/gorilla/mux v1.8.0 // indirect
)

replace example.com/sql-utils/connector => ../../internal/sql-utils/connector

replace example.com/sql-utils => ../../internal/sql-utils
