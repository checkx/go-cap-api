package domain

import (
	"api/errs"
	"api/logger"
	"database/sql"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type CustomerRepositoryDB struct {
	client *sqlx.DB
}

func NewCustomerRepositoryDB() CustomerRepositoryDB {
	connStr := "postgres://postgres:postgres@localhost/banking?sslmode=disable"
	db, err := sqlx.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return CustomerRepositoryDB{db}
}

func (d CustomerRepositoryDB) FindByID(customerID string) (*Customer, *errs.AppErr) {
	query := "select * from customers where customer_id = $1"
	var c Customer

	err := d.client.Get(&c, query, customerID)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Error("error customer data not found" + err.Error())
			return nil, errs.NewNotFoundError("customer data not found")
		} else {
			logger.Error("error scanning customer data" + err.Error())
			return nil, errs.NewUnexpectedError("unexpected database error")
		}
	}
	return &c, nil
}

func (d CustomerRepositoryDB) FindAll(status string) ([]Customer, *errs.AppErr) {
	var custs []Customer
	if status == "" {
		query := "select * from customers"
		err := d.client.Select(&custs, query)
		if err != nil {
			logger.Error("Error query customer table" + err.Error())
			return nil, errs.NewUnexpectedError("unexpected database error")
		}
	} else {
		if status == "active" {
			status = "1"
		} else {
			status = "0"
		}
		query := "select * from customers where status = $1"
		err := d.client.Select(&custs, query, status)
		if err != nil {
			logger.Error("Error query customer table" + err.Error())
			return nil, errs.NewUnexpectedError("unexpected database error")
		}
	}
	return custs, nil
}
