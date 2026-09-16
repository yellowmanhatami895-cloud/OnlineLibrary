package main

import (
	"fmt"
	"strconv"
)

func deleteOwnedBook(bookID int, userID int) {
	db.Exec("DELETE FROM books WHERE id = ? AND author_id = ?", bookID, userID)
}

func EditOwnedBooks(userID int, bookID int) {
	printBooksByID(bookID)
	record := getInputs([]string{"Enter new title", "Enter new description", "Enter new price", "Enter new category_id"})
	price, err := strconv.Atoi(record[2])
	if err != nil {
		fmt.Println(err)
		return
	}

	categoryID, err := strconv.Atoi(record[3])
	if err != nil {
		fmt.Println(err)
		return
	}
	db.Exec("UPDATE books SET title = ?,description = ?,price = ?,category_id = ? WHERE id = ? AND author_id = ?", record[0], record[1], price, categoryID, bookID, userID)
}
func printSales(userID int) {
	var bookIDs []int
	var totalSales int
	var price int
	bookSalesMap := make(map[int]int)

	row, err := db.Query("SELECT id FROM books WHERE author_id = ?", userID)
	if err != nil {
		fmt.Println(err)
		return
	}
	var bookID int
	for row.Next() {
		row.Scan(&bookID)
		bookIDs = append(bookIDs, bookID)
	}
	for _, i := range bookIDs {
		var bookSales int
		var bookName string
		row, err := db.Query("SELECT price FROM order_items WHERE book_id = ?", i)
		if err != nil {
			fmt.Println(err)
			continue
		}
		for row.Next() {
			row.Scan(&price)
			totalSales = totalSales + price
			bookSales = bookSales + price
		}
		bookSalesMap[i] = bookSales
		db.QueryRow("SELECt title FROM books WHERE id = ?", i).Scan(&bookName)
		fmt.Printf("----[Name : %s Sales : %d ]----\n", bookName, bookSales)
	}
	fmt.Printf("----------Total sales : %d----------\n", totalSales)

}
