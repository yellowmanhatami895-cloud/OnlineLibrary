package main

import "fmt"

func insertIntoUsers(record []string) {
	if len(record) != 4 {
		fmt.Println("wrong length")
		return
	}
	_, err := db.Exec("INSERT INTO users(name,email,password,role) VALUES(?,?,?,?)", record[0], record[1], record[2], record[3])
	if err != nil {
		fmt.Println(err)
		return
	}
	var id int
	db.QueryRow("SELECT id FROM users WHERE email = ?", record[1]).Scan(&id)
	fmt.Println("Your id  :", id)
}
func PasswordISCorrectUsers(userInfo []string) bool {
	var password string
	var un string
	var r string
	db.QueryRow("SELECT password,name,role FROM users WHERE id = ?", userInfo[0]).Scan(&password, &un, &r)
	if password == "" {
		fmt.Println("id does not Exists")
		return false
	}
	if password == userInfo[1] {
		username = un
		role = r
		return true
	}
	fmt.Println("wrong password")
	return false
}
