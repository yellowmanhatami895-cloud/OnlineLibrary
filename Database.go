package main

import (
	"log"

	_ "github.com/glebarez/go-sqlite"
)

func createTables() {
	db.Exec("PRAGMA foreign_keys = ON")
	db.Exec("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT NOT NULL, email TEXT NOT NULL UNIQUE, password TEXT NOT NULL, role TEXT NOT NULL DEFAULT 'Customer')")
	db.Exec("CREATE TABLE IF NOT EXISTS categories (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT NOT NULL UNIQUE)")
	db.Exec("CREATE TABLE IF NOT EXISTS books (id INTEGER PRIMARY KEY AUTOINCREMENT, title TEXT NOT NULL, description TEXT, price INTEGER NOT NULL, author_id INTEGER NOT NULL, category_id INTEGER, FOREIGN KEY (author_id) REFERENCES users(id), FOREIGN KEY (category_id) REFERENCES categories(id))")
	db.Exec("CREATE TABLE IF NOT EXISTS  orders(id INTEGER PRIMARY KEY AUTOINCREMENT, user_id INTEGER NOT NULL, total_price INTEGER NOT NULL, status TEXT NOT NULL DEFAULT 'paid', created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, FOREIGN KEY (user_id) REFERENCES users(id))")
	db.Exec("CREATE TABLE IF NOT EXISTS order_items (id INTEGER PRIMARY KEY AUTOINCREMENT, order_id INTEGER NOT NULL, book_id INTEGER NOT NULL, price INTEGER NOT NULL, FOREIGN KEY (order_id) REFERENCES orders(id), FOREIGN KEY (book_id) REFERENCES books(id))")
}
func DataBaseIsEmpty() bool {

	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM users").Scan(&count)

	if err != nil {
		log.Fatal(err)
	}
	if count == 0 {
		return true
	}
	return false
}
