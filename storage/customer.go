package storage

import (
	"database/sql"
)

type Customer struct {
	ID        int     `json:"id"`
	Name      string  `json:"customer_name"`
	Balance   float64 `json:"balance"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
	DeletedAt string  `json:"deleted_at"`
}

type ResponseError struct {
	Error interface{} `json:"error"`
	Code  int         `json:"code"`
}

func CreateCustomer(db *sql.DB, customer Customer) (Customer, error) {
	var createdCustomer Customer
	err := db.QueryRow("INSERT INTO tbl_customer (customer_name, balance) VALUES ($1, $2) RETURNING id, customer_name, balance, created_at, updated_at", customer.Name, customer.Balance).
		Scan(&createdCustomer.ID, &createdCustomer.Name, &createdCustomer.Balance, &createdCustomer.CreatedAt, &createdCustomer.UpdatedAt)
	if err != nil {
		return Customer{}, err
	}
	return createdCustomer, nil
}

func UpdateCustomer(db *sql.DB, customer Customer) (Customer, error) {
	_, err := db.Exec("UPDATE tbl_customer SET customer_name = $1, balance = $2 WHERE id = $3", customer.Name, customer.Balance, customer.ID)
	if err != nil {
		return Customer{}, err
	}
	return customer, nil
}

func DeleteCustomer(db *sql.DB, id int) (Customer, error) {
	var deletedCustomer Customer
	err := db.QueryRow("DELETE FROM tbl_customer WHERE id = $1 RETURNING id, customer_name, balance, created_at, updated_at", id).
		Scan(&deletedCustomer.ID, &deletedCustomer.Name, &deletedCustomer.Balance, &deletedCustomer.CreatedAt, &deletedCustomer.UpdatedAt)
	if err != nil {
		return Customer{}, err
	}
	return deletedCustomer, nil
}

func GetCustomer(db *sql.DB, id int) (Customer, error) {
	var customer Customer
	err := db.QueryRow("SELECT id, customer_name, balance, created_at, updated_at FROM tbl_customer WHERE id = $1", id).
		Scan(&customer.ID, &customer.Name, &customer.Balance, &customer.CreatedAt, &customer.UpdatedAt)
	if err != nil {
		return Customer{}, err
	}
	return customer, nil
}

func GetCustomers(db *sql.DB) ([]Customer, error) {
	rows, err := db.Query("SELECT id, customer_name, balance, created_at, updated_at FROM tbl_customer WHERE deleted_at IS NULL")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var customers []Customer

	for rows.Next() {
		var customer Customer
		err := rows.Scan(&customer.ID, &customer.Name, &customer.Balance, &customer.CreatedAt, &customer.UpdatedAt)
		if err != nil {
			return nil, err
		}
		customers = append(customers, customer)
	}

	return customers, nil
}
