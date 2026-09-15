package main

import (
	"fmt"
)

func printBooks() {
	row, err := db.Query("SELECT id,title,description,price,author_id,category_id FROM books")
	if err != nil {
		fmt.Println(err)
	}
	var description string
	var title string
	var price int
	var id int
	var authorID int
	var categoryID int
	var authorName string
	var category string
	for row.Next() {
		row.Scan(&id, &title, &description, &price, &authorID, &categoryID)
		db.QueryRow("SELECT name FROM users WHERE id = ? AND role = ?", authorID, "Author").Scan(&authorName)
		db.QueryRow("SELECT name FROM categories WHERE id = ?", categoryID).Scan(&category)
		fmt.Printf("----[ID : %d TITLE : %s  AUTHOR : %s  CATEGORY : %s description:%s price:%d]----\n", id, title, authorName, category, description, price)

	}
}
func printBooksByID(id int) {
	var title string
	var description string
	var categoryID int
	var authorID int
	var category string
	var author string
	var price int
	db.QueryRow("SELECT id,title,description,price,author_id,category_id FROM books WHERE id = ?", id).Scan(&id, &title, &description, &price, &authorID, &categoryID)
	db.QueryRow("SELECT name FROM users WHERE id = ?", authorID).Scan(&author)
	db.QueryRow("SELECT name FROM categories WHERE id = ?", categoryID).Scan(&category)
	fmt.Printf("----[ID : %d TITLE : %s  AUTHOR : %s  CATEGORY : %s description:%s price:%d]----\n", id, title, author, category, description, price)

}
func printBooksByTitle(title string) {
	var id int
	var description string
	var categoryID int
	var authorID int
	var category string
	var author string
	var price int
	row, err := db.Query("SELECT id,title,description,price,author_id,category_id FROM books WHERE title = ?", title)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer row.Close()
	for row.Next() {
		row.Scan(&id, &title, &description, &price, &authorID, &categoryID)
		db.QueryRow("SELECT name FROM users WHERE id = ?", authorID).Scan(&author)
		db.QueryRow("SELECT name FROM categories WHERE id = ?", categoryID).Scan(&category)
		fmt.Printf("----[ID : %d TITLE : %s  AUTHOR : %s  CATEGORY : %s description:%s price:%d]----\n", id, title, author, category, description, price)
	}

}
func printBooksByCategory(category string) {
	var id int
	var description string
	var categoryID int
	var authorID int
	var title string
	var author string
	var price int
	db.QueryRow("SELECT id FROM categories WHERE name = ?", category).Scan(&categoryID)

	row, err := db.Query("SELECT id,title,description,price,author_id,category_id FROM books WHERE category_id = ?", categoryID)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer row.Close()
	for row.Next() {
		row.Scan(&id, &title, &description, &price, &authorID, &categoryID)
		db.QueryRow("SELECT name FROM users WHERE id = ?", authorID).Scan(&author)
		fmt.Printf("----[ID : %d TITLE : %s  AUTHOR : %s  CATEGORY : %s description:%s price:%d]----\n", id, title, author, category, description, price)
	}

}

func printBooksByAuthor(author string) {
	var id int
	var description string
	var categoryID int
	var authorID int
	var title string
	var category string
	var price int
	db.QueryRow("SELECT id FROM users WHERE name = ?", author).Scan(&authorID)
	row, err := db.Query("SELECT id,title,description,price,author_id,category_id FROM books WHERE author_id= ?", authorID)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer row.Close()
	for row.Next() {
		row.Scan(&id, &title, &description, &price, &authorID, &categoryID)
		db.QueryRow("SELECT name FROM categories WHERE id = ?", categoryID).Scan(&category)
		fmt.Printf("----[ID : %d TITLE : %s  AUTHOR : %s  CATEGORY : %s description:%s price:%d]----\n", id, title, author, category, description, price)
	}

}
func printOwnedBooks(userID int) {
	var bookID int
	var orderID int
	row, err := db.Query("SELECT id FROM orders WHERE user_id = ? AND status = 'paid'", userID)
	if err != nil {
		fmt.Println(err)
		return
	}

	for row.Next() {
		row.Scan(&orderID)
		rowBook, err := db.Query("SELECT book_id FROM order_items WHERE order_id = ?", orderID)
		if err != nil {
			fmt.Println(err)
			continue
		}
		for rowBook.Next() {
			rowBook.Scan(&bookID)
			printBooksByID(bookID)
		}

	}
}

func deleteOwnedBooks(userID int, bookID int) {
	var orderID int
	var orderIDs []int
	row, err := db.Query("SELECT id FROM orders WHERE user_id = ? AND status = 'paid'", userID)
	if err != nil {
		fmt.Println("Query orders:", err)
		return
	}
	for row.Next() {
		row.Scan(&orderID)
		orderIDs = append(orderIDs, orderID)
	}
	row.Close()
	for _, id := range orderIDs {
		_, err := db.Exec("DELETE FROM order_items WHERE order_id = ? AND book_id = ?", id, bookID)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
