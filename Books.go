package main

import (
	"database/sql"
	"fmt"
	"strconv"
)

func BooksRowsPrinter(row *sql.Rows) {
	var bookID int
	var bookTitle string
	var bookDescription string
	var categoryName string
	var authorName string
	var bookPrice int
	for row.Next() {
		row.Scan(&bookID, &bookTitle, &bookPrice, &bookDescription, &authorName, &categoryName)
		fmt.Printf("[ID : %d]--[Name : %s ]--[Price : %d]--[Descriptions : %s]--[Author : %s]--[category : %s]---\n", bookID, bookTitle, bookPrice, bookDescription, authorName, categoryName)
	}
	if row.Err() != nil {
		fmt.Println(err)
		return
	}
}

func BooksRowPrinter(row *sql.Row) {
	var bookID int
	var bookTitle string
	var bookDescription string
	var categoryName string
	var authorName string
	var bookPrice int
	row.Scan(&bookID, &bookTitle, &bookPrice, &bookDescription, &authorName, &categoryName)
	fmt.Printf("[ID : %d]--[Name : %s ]--[Price : %d]--[Descriptions : %s]--[Author : %s]--[category : %s]---\n", bookID, bookTitle, bookPrice, bookDescription, authorName, categoryName)

}
func printBooks() {
	row, err := db.Query("SELECT books.id,books.title,books.price,books.description,users.name,categories.name FROM books INNER JOIN users ON books.author_id = users.id INNER JOIN categories ON books.category_id = categories.id WHERE users.role = 'Author';")
	if err != nil {
		fmt.Println(err)
		return
	}
	BooksRowsPrinter(row)
}

func printBooksByID(bookID int) {
	row := db.QueryRow("SELECT books.id,books.title,books.price,books.description,users.name,categories.name FROM books  INNER JOIN users ON books.author_id = users.id AND users.role = 'Author' INNER JOIN categories ON books.category_id = categories.id WHERE books.id = ?", bookID)
	BooksRowPrinter(row)
}
func printBooksByTitle(bookTitle string) {
	row, err := db.Query("SELECT books.id,books.title,books.price,books.description,users.name,categories.name FROM books  INNER JOIN users ON books.author_id = users.id AND users.role = 'Author' INNER JOIN categories ON books.category_id = categories.id WHERE books.title LIKE '%?%'", bookTitle)
	if err != nil {
		fmt.Println(err)
		return
	}
	BooksRowsPrinter(row)
}

func printBooksByCategory(categoryName string) {
	row, err := db.Query("SELECT books.id,books.title,books.price,books.description,users.name,categories.name FROM books  INNER JOIN users ON books.author_id = users.id AND users.role = 'Author' INNER JOIN categories ON books.category_id = categories.id WHERE categories.name LIKE '%?%'", categoryName)
	if err != nil {
		fmt.Println(err)
		return
	}
	BooksRowsPrinter(row)
}

func printBooksByAuthor(authorName string) {
	row, err := db.Query("SELECT books.id,books.title,books.price,books.description,users.name,categories.name FROM books  INNER JOIN users ON books.author_id = users.id AND users.role = 'Author' INNER JOIN categories ON books.category_id = categories.id WHERE users.name LIKE '%?%' AND users.status = 'Author'", authorName)
	if err != nil {
		fmt.Println(err)
		return
	}
	BooksRowsPrinter(row)
}
func printPurchasedBooks(userID int) {
	row, err := db.Query("SELECT  books.id,books.title,books.price,books.description,users.name,categories.name FROM books INNER JOIN order_items ON books.id = order_items.book_id INNER JOIN users ON users.id = books.author_id AND users.role = 'Author'INNER JOIN categories ON books.category_id = categories.id  INNER JOIN  orders ON  orders.id = order_items.order_id AND ? = orders.user_id AND orders.status = 'paid'", userID)
	if err != nil {
		fmt.Println(err)
	}
	BooksRowsPrinter(row)
}

func deletePurchasedBook(userID int, bookID int) {
	var booksName string
	_, err := db.Exec("DELETE FROM order_items WHERE book_id = ? AND order_id IN (SELECT id FROM orders WHERE orders.user_id = ? AND orders.status = 'paid')", bookID, userID)
	if err != nil {
		fmt.Println(err)
		return
	}
	db.QueryRow("SELECT title FROM books WHERE id = ?", bookID).Scan(&booksName)
	fmt.Println(booksName + " Deleted")
}
func getBookRow() (i int, title string, des string, price int, authorID int, categoryID int) {
	row := getInputs([]string{"Enter title", "Enter description", "Enter price", "Enter author id", "Enter category id"})
	for _, r := range row {
		if r == "0" {

		}
	}
	if i == 0 {
		return 0, "", "", 0, 0, 0
	}
	price, err = strconv.Atoi(row[2])
	if err != nil {
		fmt.Println(err)
		return 0, "", "", 0, 0, 0
	}
	authorID, err = strconv.Atoi(row[3])
	if err != nil {
		fmt.Println(err)
		return 0, "", "", 0, 0, 0
	}
	categoryID, err = strconv.Atoi(row[4])
	if err != nil {
		fmt.Println(err)
		return 0, "", "", 0, 0, 0
	}
	return 1, row[0], row[1], price, authorID, categoryID
}
func insertIntoBooks() int {
	i, title, descriptions, price, authorID, categoryID := getBookRow()
	if i == 0 {
		return 0
	}
	_, err = db.Exec("INSERT INTO books(title,description,price,author_id,category_id) VALUES(?,?,?,?,?)", title, descriptions, price, authorID, categoryID)
	if err != nil {
		fmt.Println("Insert book:", err)
		return 0
	}
	return 1
}
func deleteBooks(bookID int) {
	var bookName string
	db.QueryRow("SELECT name FROM books WHERE id = ?").Scan(&bookName)
	db.Exec("DELETE FROM books WHERE id =?", bookID)
	fmt.Println(bookName + "  DELETED")
}
func editBook(bookID int) {
	printBooksByID(bookID)
	fmt.Println("------/Enter new row/-----")
	i, title, descriptions, price, authorID, categoryID := getBookRow()
	if i == 0 {
		return
	}
	db.Exec("UPDATE books SET title = ?, description = ?,price = ?,author_id = ?,category_id = ? WHERE id = ?", title, descriptions, price, authorID, categoryID, bookID)
}
