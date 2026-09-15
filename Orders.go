package main

import (
	"fmt"
)

func insertBooksIntCart(bookIDs []int, userID int, status string) {

	var orderStatus string
	var orderID int
	var oldTotalPrice int
	var totalPrice int
	db.QueryRow("SELECT status, id, total_price FROM orders WHERE user_id = ? AND status = 'in cart'", userID).Scan(&orderStatus, &orderID, &oldTotalPrice)

	if status == "in cart" {
		for _, bookID := range bookIDs {
			var checkID int
			var price int
			db.QueryRow("SELECT id FROM order_items WHERE book_id = ? AND order_id = ?", bookID, orderID).Scan(&checkID)
			if checkID != 0 {
				fmt.Println("book already exists, order item id:", checkID)
				continue
			}
			db.QueryRow("SELECT price FROM books WHERE id = ?", bookID).Scan(&price)

			result, _ := db.Exec("INSERT INTO order_items(order_id, book_id, price) VALUES(?, ?, ?)", orderID, bookID, price)

			orderItemID, _ := result.LastInsertId()

			totalPrice = totalPrice + price

			fmt.Println("Order item ID:", orderItemID)
		}

		db.Exec("UPDATE orders SET total_price = ? WHERE id = ?", oldTotalPrice+totalPrice, orderID)
		return
	}

	result, _ := db.Exec("INSERT INTO orders(user_id, status, total_price) VALUES(?, ?, ?)", userID, status, 0)

	orderID2, _ := result.LastInsertId()

	fmt.Println("Order ID:", orderID2)

	for _, bookID := range bookIDs {

		var price int

		db.QueryRow("SELECT price FROM books WHERE id = ?", bookID).Scan(&price)

		result, _ := db.Exec("INSERT INTO order_items(order_id, book_id, price) VALUES(?, ?, ?)", orderID2, bookID, price)

		orderItemID, _ := result.LastInsertId()

		totalPrice += price
		fmt.Println("Order item ID:", orderItemID)
	}

	db.Exec("UPDATE orders SET total_price = ? WHERE id = ?", totalPrice, orderID2)
}
func DeleteBookFromOrderItems(bookID int, userID int, status string) {
	var price int
	var totalPrice int
	var orderID int
	db.QueryRow("SELECT id FROM orders WHERE user_id = ? AND status = ?", userID, status).Scan(&orderID)
	db.QueryRow("SELECT price FROM books WHERE id = ?", bookID).Scan(&price)
	if price == 0 {
		fmt.Println("book does not Exist")
		return
	}
	db.Exec("DELETE FROM order_items WHERE order_id = ? AND book_id = ?", orderID, bookID)
	db.QueryRow("SELECT total_price FROM orders WHERE id = ?", orderID).Scan(&totalPrice)
	db.Exec("UPDATE orders set total_price = ? WHERE id = ?", totalPrice-price, orderID)
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
func payCart(userID int) {
	var price int
	var orderID int
	db.QueryRow("SELECT price,id FROM orders WHERE status = 'in cart' AND userID = ?", userID).Scan(&price, &orderID)
	fmt.Println("price :", price)
	input := getIntInput("Enter 1 to pay")
	if input == 1 {
		db.Exec("UPDATE orders SET status = 'paid' WHERE id = ?", orderID)
	}

}
