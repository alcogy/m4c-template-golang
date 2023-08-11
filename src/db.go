package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Customer struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	Zip             string `json:"zip"`
	Address         string `json:"address"`
	Tel             string `json:"tel"`
	Email           string `json:"email"`
}

func connection() *sql.DB {
	db, err := sql.Open("postgres", "host=m4c-database port=5432 user=user password=pass dbname=m4c-db sslmode=disable")

	if err != nil {
		fmt.Println(err)
	}

	return db
}

func SelectAllCustomers() []Customer {
	db := connection()
	defer db.Close()
	// SELECT
	sql := "SELECT * FROM customer order by id"
	rows, err := db.Query(sql)

	if err != nil {
		log.Fatal(err)
	}

	var customers []Customer
	for rows.Next() {
		var c Customer
		rows.Scan(
			&c.ID,
			&c.Name,
			&c.Zip,
			&c.Address,
			&c.Tel,
			&c.Email)
		customers = append(customers, c)
	}

	return customers
}

func SelectCustomerById(id string) (Customer, error) {
	db := connection()
	defer db.Close()

	stmt, err := db.Prepare("SELECT * FROM customer WHERE id = $1")
	if err != nil {
		log.Fatal(err)
	}

	var c Customer
	err = stmt.QueryRow(id).Scan(
		&c.ID,
		&c.Name,
		&c.Zip,
		&c.Address,
		&c.Tel,
		&c.Email)

	if err != nil {
		return c, err
	}

	return c, nil
}

func InsertCustomer(customer Customer) {
	db := connection()
	defer db.Close()

	_, err := db.Exec("INSERT INTO customer VALUES (DEFAULT, $1, $2, $3, $4, $5);", customer.Name, customer.Zip, customer.Address, customer.Tel, customer.Email)

	if err != nil {
		log.Fatal(err)
	}
}

func updateCustomer(customer Customer) {
	db := connection()
	defer db.Close()

	_, err := db.Exec("UPDATE customer SET name = $1, zip = $2, address = $3, tel = $4, email = $5 WHERE id = $6;", customer.Name, customer.Zip, customer.Address, customer.Tel, customer.Email, customer.ID)

	if err != nil {
		log.Fatal(err)
	}

}

func deleteCustomer(id int) {
	db := connection()
	defer db.Close()

	_, err := db.Exec("DELETE FROM customer WHERE id = $1;", id)

	if err != nil {
		log.Fatal(err)
	}
}