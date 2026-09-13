package main

import "fmt"

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
func insertBookIntoOrders(bookIDs []int, userID int, status string) {
	result, err := db.Exec("INSERT INTO orders(user_id, status,total_price) VALUES(?, ?,?)", userID, status, 0)
	if err != nil {
		fmt.Println(err)
		return
	}

	orderID, err := result.LastInsertId()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Order ID:", orderID)
	var price int
	var totalPrice int
	for _, i := range bookIDs {
		db.QueryRow("SELECT price FROM books WHERE id = ?", i).Scan(&price)
		_, err := db.Exec("INSERT INTO order_items(order_id,book_id,price) VALUES(?,?,?)", orderID, i, price)

		if err != nil {
			fmt.Println(err)
			return
		}
		totalPrice = totalPrice + price
	}
	db.Exec("UPDATE orders set total_price = ? WHERE id = ?", totalPrice, orderID)
}
