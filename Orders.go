package main

import "fmt"

func insertBooksIntoOrders(bookIDs []int, userID int, status string) {
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
		result, err := db.Exec("INSERT INTO order_items(order_id,book_id,price) VALUES(?,?,?)", orderID, i, price)
		order_itemID, err := result.LastInsertId()
		if err != nil {
			fmt.Println(err)
			return
		}
		if err != nil {
			fmt.Println(err)
			return
		}
		totalPrice = totalPrice + price
		fmt.Println("Order item ID:", order_itemID)
	}
	db.Exec("UPDATE orders set total_price = ? WHERE id = ?", totalPrice, orderID)
}
func DeleteBookFromOrderItems(bookID int, orderID int, status string) {
	db.Exec("DELETE FROM order_items WHERE order_id = ? AND book_id = ?", orderID, bookID)
}
func printOrders(userID int) {
	row, err := db.Query("SELECT * FROM orders WHERE user_id = ?", userID)
	if err != nil {
		fmt.Println(err)
	}
	var orderID int
	var totalPrice int
	var status string
	var time string
	for row.Next() {
		row.Scan(&orderID, &userID, &totalPrice, &status, &time)
		fmt.Printf("--[order id : %d  user id : %d total price : %d status : %s time : %s]--\n", orderID, userID, totalPrice, status, time)
	}
}
