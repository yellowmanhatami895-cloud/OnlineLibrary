package main

import (
	"fmt"
)

func deleteOwnedBook(bookID int, userID int) {
	db.Exec("DELETE FROM books WHERE id = ? AND author_id = ?", bookID, userID)
}

func EditOwnedBooks(userID int, bookID int) {
	printBooksByID(bookID)
	fmt.Println("**** Enter new row ****")
	i, title, description, price, authorID, categoryID := getBookRow()
	if i == 0 {
		return
	}
	db.Exec("UPDATE books SET title = ?,description = ?,price = ?,category_id = ? WHERE id = ? AND author_id = ?", title, description, price, authorID, categoryID)
}
func printSales(userID int) {
	sales, err := db.Query("SELECT order_items.price,books.title FROM order_items INNER JOIN books ON books.id = order_items.book_id INNER JOIN orders ON order_items.order_id = orders.id WHERE books.author_id = ? AND orders.status = 'paid'", userID)
	if err != nil {
		fmt.Println(err)
		return
	}
	var totalPrice int
	var bookTitle string
	for sales.Next() {
		var price int
		sales.Scan(&price, &bookTitle)
		totalPrice = totalPrice + price
		fmt.Printf("---Name : %s  | Price : %d ---\n", bookTitle, price)
	}
	err = sales.Err()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf(" ------ Total Sales %d ------\n", totalPrice)
}
