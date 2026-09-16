package main

import (
	"fmt"
	"strconv"
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
		description, title, price, id, authorID, categoryID, authorName, category = "", "", 0, 0, 0, 0, "", ""
		row.Scan(&id, &title, &description, &price, &authorID, &categoryID)
		db.QueryRow("SELECT name FROM users WHERE id = ? AND role = ?", authorID, "Author").Scan(&authorName)
		db.QueryRow("SELECT name FROM categories WHERE id = ?", categoryID).Scan(&category)
		fmt.Printf("----[ID : %d TITLE : %s  AUTHOR : %s  CATEGORY : %s description:%s price:%d]----\n", id, title, authorName, category, description, price)

	}
}
func printBooksByID(id int) {
	var description string
	var title string
	var price int
	var authorID int
	var categoryID int
	var authorName string
	var category string

	err := db.QueryRow("SELECT title,description,price,author_id,category_id FROM books WHERE id = ?", id).Scan(&title, &description, &price, &authorID, &categoryID)
	if err != nil {
		fmt.Println("Book:", err)
		return
	}

	err = db.QueryRow("SELECT name FROM users WHERE id = ? AND role = ?", authorID, "Author").Scan(&authorName)
	fmt.Println("authorID:", authorID)
	if err != nil {
		fmt.Println("Author:", err)
		return
	}

	err = db.QueryRow("SELECT name FROM categories WHERE id = ?", categoryID).Scan(&category)
	if err != nil {
		fmt.Println("Category:", err)
		return
	}

	fmt.Printf("----[ID : %d TITLE : %s AUTHOR : %s CATEGORY : %s description:%s price:%d]----\n", id, title, authorName, category, description, price)
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

func deletePurchasedBook(userID int, bookID int) {
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
func getBookRow() (int, []string) {
	row := getInputs([]string{"Enter title", "Enter description", "Enter price", "Enter author id", "Enter category id"})
	for _, r := range row {
		if r == "0" {
			return 0, []string{}
		}
	}
	return 1, row
}
func insertIntoBooks() int {
	i, row := getBookRow()
	if i == 0 {
		return 0
	}
	price, err := strconv.Atoi(row[2])
	if err != nil {
		fmt.Println(err)
		return 0
	}
	authorID, err := strconv.Atoi(row[3])
	if err != nil {
		fmt.Println(err)
		return 0
	}
	categoryID, err := strconv.Atoi(row[4])
	if err != nil {
		fmt.Println(err)
		return 0
	}
	_, err = db.Exec("INSERT INTO books(title,description,price,author_id,category_id) VALUES(?,?,?,?,?)", row[0], row[1], price, authorID, categoryID)
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
	i, row := getBookRow()
	if i == 0 {
		return
	}
	price, err := strconv.Atoi(row[2])
	if err != nil {
		fmt.Println(err)
		return
	}
	authorID, err := strconv.Atoi(row[3])
	if err != nil {
		fmt.Println(err)
		return
	}
	categoryID, err := strconv.Atoi(row[4])
	if err != nil {
		fmt.Println(err)
		return
	}
	db.Exec("UPDATE books SET title = ?, description = ?,price = ?,author_id = ?,category_id = ? WHERE id = ?", row[0], row[1], price, authorID, categoryID, bookID)
}
