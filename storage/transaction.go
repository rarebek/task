package storage

import (
	"database/sql"
	"fmt"
	"time"
)

type Transaction struct {
	ID         int
	CustomerID int
	ItemID     int
	Qty        int
	Amount     float64
	CreatedAt  string
	UpdatedAt  string
	DeletedAt  string
}

type TransactionView struct {
	ID           int        `json:"id"`
	CustomerID   int        `json:"customer_id"`
	CustomerName string     `json:"customer_name"`
	ItemID       int        `json:"item_id"`
	ItemName     string     `json:"item_name"`
	Qty          int        `json:"qty"`
	Price        float64    `json:"price"`
	Amount       float64    `json:"amount"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at"`
}

func GetTransactions(db *sql.DB) ([]Transaction, error) {
	rows, err := db.Query("SELECT id, customer_id, item_id, qty, amount, created_at, updated_at FROM tbl_transaction")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []Transaction

	for rows.Next() {
		var transaction Transaction
		err := rows.Scan(&transaction.ID, &transaction.CustomerID, &transaction.ItemID, &transaction.Qty, &transaction.Amount, &transaction.CreatedAt, &transaction.UpdatedAt)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}

	return transactions, nil
}

func CreateTransaction(db *sql.DB, transaction Transaction) (Transaction, error) {
	var createdTransaction Transaction
	err := db.QueryRow("INSERT INTO tbl_transaction (customer_id, item_id, qty, amount) VALUES ($1, $2, $3, $4) RETURNING id, customer_id, item_id, qty, amount, created_at, updated_at",
		transaction.CustomerID, transaction.ItemID, transaction.Qty, transaction.Amount).
		Scan(&createdTransaction.ID, &createdTransaction.CustomerID, &createdTransaction.ItemID, &createdTransaction.Qty, &createdTransaction.Amount, &createdTransaction.CreatedAt, &createdTransaction.UpdatedAt)
	if err != nil {
		return Transaction{}, err
	}
	return createdTransaction, nil
}

func UpdateTransaction(db *sql.DB, transaction Transaction) (Transaction, error) {
	_, err := db.Exec("UPDATE tbl_transaction SET qty = $1, amount = $2, updated_at = CURRENT_TIMESTAMP WHERE id = $3", transaction.Qty, transaction.Amount, transaction.ID)
	if err != nil {
		return Transaction{}, err
	}
	return transaction, nil
}

func DeleteTransaction(db *sql.DB, id int) error {
	_, err := db.Exec("UPDATE tbl_transaction SET deleted_at = CURRENT_TIMESTAMP WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func GetTransaction(db *sql.DB, id int) (Transaction, error) {
	var transaction Transaction
	err := db.QueryRow("SELECT id, customer_id, item_id, qty, amount, created_at, updated_at FROM tbl_transaction WHERE id = $1", id).
		Scan(&transaction.ID, &transaction.CustomerID, &transaction.ItemID, &transaction.Qty, &transaction.Amount, &transaction.CreatedAt, &transaction.UpdatedAt)
	if err != nil {
		return Transaction{}, err
	}
	return transaction, nil
}

func GetTransactionDetailsWithCustomerAndItem(db *sql.DB) ([]TransactionView, error) {
	rows, err := db.Query("SELECT tv.id, tv.customer_id, c.customer_name, tv.item_id, i.item_name, tv.qty, tv.price, tv.amount, tv.created_at, tv.updated_at FROM TransactionViews tv INNER JOIN tbl_customer c ON tv.customer_id = c.id INNER JOIN tbl_items i ON tv.item_id = i.id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []TransactionView

	for rows.Next() {
		var transaction TransactionView
		if err := rows.Scan(&transaction.ID, &transaction.CustomerID, &transaction.CustomerName, &transaction.ItemID, &transaction.ItemName, &transaction.Qty, &transaction.Price, &transaction.Amount, &transaction.CreatedAt, &transaction.UpdatedAt); err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return transactions, nil
}

func FilterTransactions(db *sql.DB, id int, customerName, itemName string) ([]TransactionView, error) {
	query := "SELECT tv.id, tv.customer_id, c.customer_name, tv.item_id, i.item_name, tv.qty, tv.price, tv.amount, tv.created_at, tv.updated_at FROM TransactionViews tv INNER JOIN tbl_customer c ON tv.customer_id = c.id INNER JOIN tbl_items i ON tv.item_id = i.id WHERE true"

	if id != 0 {
		query += fmt.Sprintf(" AND tv.id = %d", id)
	}

	if customerName != "" {
		query += fmt.Sprintf(" AND c.customer_name = '%s'", customerName)
	}

	if itemName != "" {
		query += fmt.Sprintf(" AND i.item_name = '%s'", itemName)
	}

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []TransactionView

	for rows.Next() {
		var transaction TransactionView
		if err := rows.Scan(&transaction.ID, &transaction.CustomerID, &transaction.CustomerName, &transaction.ItemID, &transaction.ItemName, &transaction.Qty, &transaction.Price, &transaction.Amount, &transaction.CreatedAt, &transaction.UpdatedAt); err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return transactions, nil
}
