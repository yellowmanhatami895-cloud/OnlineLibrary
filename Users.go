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
	fmt.Println("Wrong password")
	return false
}
func editProfile(userID int) {
	var oldEmail string
	printUsersByID(userID)
	info := getInputs([]string{"Enter new name", "Enter new email", "Enter new password"})
	db.QueryRow("SELECT email FROM users WHERE id = ?").Scan(&oldEmail)
	row, err := db.Query("SELECT email FROM users")
	if err != nil {
		fmt.Println(err)
		return
	}
	var email string
	for row.Next() {
		row.Scan(&email)
		if oldEmail == email {
			continue
		}
		if email == info[1] {
			fmt.Println("email is already used")
			return
		}
	}
	db.Exec("UPDATE users SET name = ?, email = ?, password = ? WHERE id = ?", info[0], info[1], info[2], userID)
}
func printUsersByID(userID int) {
	var name string
	var email string
	var password string
	var role string
	db.QueryRow("SELECT * FROM users WHERE id = ? ", userID).Scan(&userID, &name, &email, &password, &role)
	fmt.Printf("---[Id : %d name : %s email : %s password : %s role : %s]--\n", userID, name, email, password, role)
}
