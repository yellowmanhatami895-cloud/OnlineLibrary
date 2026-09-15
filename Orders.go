package main

import (
	"database/sql"
	"fmt"
)

func insertBooksIntCart(bookIDs []int, userID int, status string) {

	var orderStatus string
	var orderID int
	var oldTotalPrice int
	var totalPrice int
	err := db.QueryRow("SELECT status, id, total_price FROM orders WHERE user_id = ? AND status = 'in cart'", userID).Scan(&orderStatus, &orderID, &oldTotalPrice)
	if err == sql.ErrNoRows {
		db.Exec("INSERT INTO orders(user_id,total_price,status) VALUES(?,?,?)", userID, 0, "in cart")
		err := db.QueryRow("SELECT status, id, total_price FROM orders WHERE user_id = ? AND status = 'in cart'", userID).Scan(&orderStatus, &orderID, &oldTotalPrice)
		if err != nil {
			fmt.Println(err)
		}
	}
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

			result, err := db.Exec("INSERT INTO order_items(order_id, book_id, price) VALUES(?, ?, ?)", orderID, bookID, price)

			if err != nil {
				fmt.Println(err)
				return
			}

			orderItemID, err := result.LastInsertId()

			if err != nil {
				fmt.Println(err)
				return
			}

			totalPrice = totalPrice + price

			fmt.Println("Order item ID:", orderItemID)
		}

		db.Exec("UPDATE orders SET total_price = ? WHERE id = ?", oldTotalPrice+totalPrice, orderID)
		return
	}

	result, err := db.Exec("INSERT INTO orders(user_id, status, total_price) VALUES(?, ?, ?)", userID, status, 0)

	if err != nil {
		fmt.Println(err)
		return
	}

	orderID2, err := result.LastInsertId()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Order ID:", orderID2)

	for _, bookID := range bookIDs {

		var price int

		db.QueryRow("SELECT price FROM books WHERE id = ?", bookID).Scan(&price)

		result, err := db.Exec("INSERT INTO order_items(order_id, book_id, price) VALUES(?, ?, ?)", orderID2, bookID, price)

		if err != nil {
			fmt.Println(err)
			return
		}

		orderItemID, err := result.LastInsertId()

		if err != nil {
			fmt.Println(err)
			return
		}

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
	db.QueryRow("SELECT total_price,id FROM orders WHERE status = 'in cart' AND user_id = ?", userID).Scan(&price, &orderID)
	fmt.Println("price :", price)
	input := getIntInput("Enter 1 to pay")
	if input == 1 {
		db.Exec("UPDATE orders SET status = 'paid' WHERE id = ?", orderID)
	}

}
func deleteAllOrdersItems(userID int) {
	var orderID int
	db.QueryRow("SELECT id FROM orders WHERE user_id = ? AND status = 'in cart'", userID).Scan(&orderID)
	db.Exec("DELETE FROM order_items WHERE order_id = ?", orderID)
	db.Exec("UPDATE orders SET total_price = 0 WHERE status = 'in cart' AND user_id = ?", userID)
}
