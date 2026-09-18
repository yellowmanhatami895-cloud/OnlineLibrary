package main

import "fmt"

func checkEmail(oldEmail string, newEmail string) {
	row, err := db.Query("SELECT email FROM users")
	if err != nil {
		fmt.Println(err)
		return
	}
	var email string
	for row.Next() {
		row.Scan(&email)
		if email == oldEmail {
			continue
		}
		if email == newEmail {
			fmt.Println("Email is already used")
		}
	}
	err = row.Err()
	if err != nil {
		fmt.Println(err)
		return
	}
}
func insertCustomerIntoUsers() {
	record := getInputs([]string{"Enter name", "Email", "Password"})
	record = append(record, "Customer")
	insertIntoUsers(record)
}
func insertUsers() {
	record := getInputs([]string{"Enter name", "Email", "Password", "role"})
	insertIntoUsers(record)
}
func insertIntoUsers(record []string) {
	if len(record) != 4 {
		fmt.Println("wrong length")
		return
	}
	checkEmail("", record[1])
	result, err := db.Exec("INSERT INTO users(name,email,password,role) VALUES(?,?,?,?)", record[0], record[1], record[2], record[3])
	if err != nil {
		fmt.Println(err)
		return
	}
	id, err := result.LastInsertId()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Your id  :", id)
}
func PasswordISCorrectUsers(userInfo []string) bool {
	var password string
	var un string
	var r string
	db.QueryRow("SELECT password,name,role FROM users WHERE id = ?", userInfo[0]).Scan(&password, &un, &r)
	if password == "" {
		fmt.Println("id does not Exist")
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
	if info[0] == "0" || info[1] == "0" || info[2] == "0" {
		return
	}
	db.QueryRow("SELECT email FROM users WHERE id = ?", userID).Scan(&oldEmail)
	checkEmail(oldEmail, info[1])
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
func deleteFromUsers(userID int) {
	var name string
	db.QueryRow("SELECT name FROM users WHERE id = ?", userID).Scan(&name)
	_, err := db.Exec("DELETE FROM users WHERE id = ?", userID)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(name + " DELETED")
}

func printUsersByEmail(userEmail string) {
	var id int
	var name string
	var password string
	var role string
	var email string
	db.QueryRow("SELECT * FROM users WHERE email = ?", userEmail).Scan(&id, &name, &email, &password, &role)
	fmt.Printf("---[Id : %d name : %s email : %s password : %s role : %s]--\n", id, name, email, password, role)
}
func printUsersByName(name string) {
	var id int
	var password string
	var role string
	var email string
	row, err := db.Query("SELECT * FROM users WHERE name = ?", name)
	if err != nil {
		fmt.Println(err)
		return
	}
	for row.Next() {
		row.Scan(&id, &name, &email, &password, &role)
		fmt.Printf("---[Id : %d name : %s email : %s password : %s role : %s]--\n", id, name, email, password, role)
	}
	err = row.Err()
	if err != nil {
		fmt.Println(err)
		return
	}
}
func printUsersByRole(role string) {
	var id int
	var password string
	var name string
	var email string
	row, err := db.Query("SELECT * FROM users WHERE role = ?", role)
	if err != nil {
		fmt.Println(err)
		return
	}
	for row.Next() {
		row.Scan(&id, &name, &email, &password, &name)
		fmt.Printf("---[Id : %d name : %s email : %s password : %s role : %s]--\n", id, name, email, password, role)
	}

	err = row.Err()
	if err != nil {
		fmt.Println(err)
		return
	}
}
func printUsers() {
	var id int
	var password string
	var name string
	var email string
	var role string
	row, err := db.Query("SELECT * FROM users")
	if err != nil {
		fmt.Println(err)
		return
	}
	for row.Next() {
		row.Scan(&id, &name, &email, &password, &role)
		fmt.Printf("---[Id : %d name : %s email : %s password : %s role : %s]--\n", id, name, email, password, role)
	}

	err = row.Err()
	if err != nil {
		fmt.Println(err)
		return
	}
}
