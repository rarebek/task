package storage

import (
	"database/sql"
)

type Item struct {
	ID        int     `json:"id"`
	Name      string  `json:"item_name"`
	Cost      float64 `json:"cost"`
	Price     float64 `json:"price"`
	Sort      int     `json:"sort"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
	DeletedAt string  `json:"deleted_at"`
}

func CreateItem(db *sql.DB, item Item) (Item, error) {
	var createdItem Item
	err := db.QueryRow("INSERT INTO tbl_items (item_name, cost, price, sort) VALUES ($1, $2, $3, $4) RETURNING id, item_name, cost, price, sort, created_at", item.Name, item.Cost, item.Price, item.Sort).
		Scan(&createdItem.ID, &createdItem.Name, &createdItem.Cost, &createdItem.Price, &createdItem.Sort, &createdItem.CreatedAt)
	if err != nil {
		return Item{}, err
	}
	return createdItem, nil
}

func UpdateItem(db *sql.DB, item Item) (Item, error) {
	_, err := db.Exec("UPDATE tbl_items SET item_name = $1, cost = $2, price = $3, sort = $4 WHERE id = $5", item.Name, item.Cost, item.Price, item.Sort, item.ID)
	if err != nil {
		return Item{}, err
	}
	return item, nil
}

func DeleteItem(db *sql.DB, id int) (Item, error) {
	var deletedItem Item
	err := db.QueryRow("DELETE FROM tbl_items WHERE id = $1 RETURNING id, item_name, cost, price, sort, created_at", id).
		Scan(&deletedItem.ID, &deletedItem.Name, &deletedItem.Cost, &deletedItem.Price, &deletedItem.Sort, &deletedItem.CreatedAt)
	if err != nil {
		return Item{}, err
	}
	return deletedItem, nil
}

func GetItem(db *sql.DB, id int) (Item, error) {
	var item Item
	err := db.QueryRow("SELECT id, item_name, cost, price, sort, created_at FROM tbl_items WHERE id = $1", id).
		Scan(&item.ID, &item.Name, &item.Cost, &item.Price, &item.Sort, &item.CreatedAt)
	if err != nil {
		return Item{}, err
	}
	return item, nil
}

func GetItems(db *sql.DB) ([]Item, error) {
	rows, err := db.Query("SELECT id, item_name, cost, price, sort, created_at, updated_at FROM tbl_items WHERE deleted_at IS NULL")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []Item

	for rows.Next() {
		var item Item
		err := rows.Scan(&item.ID, &item.Name, &item.Cost, &item.Price, &item.Sort, &item.CreatedAt, &item.UpdatedAt)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}
