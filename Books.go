package main

import "fmt"

func printBooks() {
	row, err := db.Query("SELECT id,title,author_id,category_id FROM books")
	if err != nil {
		fmt.Println(err)
	}
	var title string
	var id int
	var authorID int
	var categoryID int
	var authorName string
	var category string
	for row.Next() {
		row.Scan(&id, &title, &authorID, &categoryID)
		db.QueryRow("SELECT name FROM users WHERE id = ? AND role = ?", authorID, "Author").Scan(&authorName)
		db.QueryRow("SELECT name FROM categories WHERE id = ?", categoryID).Scan(&category)
		fmt.Printf("----[ID : %d TITLE : %s  AUTHOR : %s  CATEGORY : %s]----\n", id, title, authorName, category)

	}
}
