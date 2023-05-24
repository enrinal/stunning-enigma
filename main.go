package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"rg-comp-api/api"
	"rg-comp-api/repository"
)

const (
	host     = "localhost"
	port     = 54320
	user     = "user"
	password = "admin"
	dbname   = "rg_comp"
)

// ConnectDB is a function to connect to database using postgres driver
func ConnectDB() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func Migration(db *sql.DB) error {

	// table department
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS departments (id SERIAL PRIMARY KEY, name VARCHAR(255))`)
	if err != nil {
		return err
	}

	// table employee
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS employees (id SERIAL PRIMARY KEY, name VARCHAR(255), address VARCHAR(255), salary INTEGER, department_id INTEGER REFERENCES departments(id))`)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	db, err := ConnectDB()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	err = Migration(db)
	if err != nil {
		panic(err)
	}

	emplRepo := repository.NewEmployeeRepo(db)
	mainApi := api.NewApi(emplRepo)
	mainApi.Start()
}
