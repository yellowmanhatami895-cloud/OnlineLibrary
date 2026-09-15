package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strconv"

	_ "github.com/glebarez/go-sqlite"
)

var scanner = bufio.NewScanner(os.Stdin)
var db *sql.DB
var err error
var username string
var role string

func getIntInput(title string) int {
	input := getInputs([]string{title})
	intInput, err := strconv.Atoi(input[0])
	if err != nil {
		fmt.Println(err)
		return 0
	}
	return intInput
}
func getInputs(titles []string) []string {
	var inputs []string
	for _, s := range titles {
		for {
			fmt.Printf("------%s------\n", s)
			scanner.Scan()
			if scanner.Text() == "" {
				fmt.Println("Enter Something")
				continue
			}
			break
		}
		if scanner.Text() == "0" {
			return []string{"0"}
		}
		inputs = append(inputs, scanner.Text())
	}
	return inputs
}
func getMenu(title string, subjects []string) int {
	fmt.Printf("-----%s-----\n", title)
	for i, s := range subjects {
		fmt.Printf("%d.%s\n", i+1, s)
	}
	input := getInputs([]string{"Enter menu index"})

	intInput, err := strconv.Atoi(input[0])
	if err != nil {
		fmt.Println("Enter menu index")
		return -1
	}
	if intInput < 0 || intInput > len(subjects) {
		fmt.Printf("Enter menu index between 1 to %d\n", len(subjects))
		return -1
	}
	return intInput
}
func main() {
	db, err = sql.Open("sqlite", "DataBase/OnlineLibrary.db")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
	createTables()
	if DataBaseIsEmpty() {
		fmt.Println("DataBase is Empty")
		record := getInputs([]string{"Enter first username", "Enter email", "Enter first password"})
		insertIntoUsers(append(record, "Admin"))
	}
MainLoop:
	for {
		input := getMenu("Online library", []string{"Login", "Create Account"})
		switch input {
		case 0:
			break MainLoop
		case 1:
		LoginLoop:
			for {
				userInfo := getInputs([]string{"Enter id", "Enter password"})
				userid := userInfo[0]
				id, err := strconv.Atoi(userid)
				if err != nil {
					fmt.Println(err)
					return
				}
				if userInfo[0] == "0" {
					break LoginLoop
				}
				if PasswordISCorrectUsers(userInfo) {
					switch role {
					case "Customer":
					CustomerMainLoop:
						for {
							fmt.Println("welcome Customer " + username)
							input := getMenu("Customer Panel", []string{"Purchase Book", "Edit Profile", "View Orders", "My Library", "Logout"})
							switch input {
							case 0:
								break CustomerMainLoop
							case 1:
							PurchaseLoop:
								for {
									input := getMenu("Purchase Panel", []string{"Cart", "Search books", "View all books"})
									switch input {
									case 0:
										break PurchaseLoop
									case 1:
									cartLoop:
										for {
											input := getMenu("Cart", []string{"Add to cart", "Remove from cart", "Pay", "Remove all", "View cart items"})
											switch input {
											case 0:
												break cartLoop
											case 1:
												var bookIDs []int
											AddLoop:
												for {
													bookID := getIntInput("Enter Book ids ")
													if bookID == 0 {
														break AddLoop
													}
													bookIDs = append(bookIDs, bookID)
												}
												insertBooksIntCart(bookIDs, id, "in cart")
											case 2:
											DeleteLoop:
												for {
													bookID := getIntInput("Enter book id")
													if bookID == 0 {
														break DeleteLoop
													}
													DeleteBookFromOrderItems(bookID, id, "in cart")
												}
											case 3:
												payCart(id)
											case 4:
												deleteAllOrdersItems(id)
											case 5:
												printCartItems(id)
											}
										}
									case 2:
									SearchLoop:
										for {
											input := getMenu("Search by :", []string{"id", "title", "Category", "author"})
											switch input {
											case 0:
												break SearchLoop
											case 1:
												id := getIntInput("Enter id")
												if id == 0 {
													break SearchLoop
												}
												printBooksByID(id)
											case 2:
												title := getInputs([]string{"Enter title"})
												if title[0] == "0" {
													break SearchLoop
												}
												printBooksByTitle(title[0])
											case 3:
												category := getInputs([]string{"Enter category"})
												if category[0] == "0" {
													break SearchLoop
												}
												printBooksByCategory(category[0])
											case 4:
												author := getInputs([]string{"Enter author"})
												if author[0] == "0" {
													break SearchLoop
												}
												printBooksByAuthor(author[0])
											}

										}
									case 3:
										printBooks()

										printOwnedBooks(id)
									}
								}
							case 2:
								editProfile(id)
							case 3:
								printOrders(id)
							case 4:
							OwnedBooksLoop:
								for {
									input := getMenu("Owned books", []string{"Delete", "Print"})
									switch input {
									case 0:
										break OwnedBooksLoop
									case 1:
									DeleteBookLoop:
										for {
											input := getIntInput("Enter book id")
											if input == 0 {
												break DeleteBookLoop
											}
											deleteOwnedBooks(id, input)
										}
									case 2:
										printOwnedBooks(id)
									case 7:

									}
								}
							case 5:
								userInfo[0] = ""
								userInfo[1] = ""
								id = 0
								userid = ""
								break LoginLoop
							}
						}
					case "Admin":
					StaffMainLoop:
						for {
							fmt.Println("welcome Admin " + username)
							input := getMenu("Staff panel", []string{"Manage users", "Manage books", "View orders", "View sales", "Logout"}) // add del edit print book and users
							switch input {
							case 0:
								break StaffMainLoop
							case 1:
								getMenu("Books", []string{"Add", "Delete", "Search"})
							case 2:
							BooksLoop:
								for {
									input := getMenu("Books", []string{"Add", "Delete", "Edit", "Search"})
									switch input {
									case 0:
										break BooksLoop
									case 1:
										for {
											i := insertIntoBooks()
											if i == 0 {
												break
											}
										}
									case 2:
									DeleteBooksLoop:
										for {
											id := getIntInput("Enter id")
											if id == 0 {
												break DeleteBooksLoop
											}
											deleteBooks(id)
										}
									case 3:
									EditLoop:
										for {
											input := getIntInput("Enter book id")
											if input == 0 {
												break EditLoop
											}
											editBook(input)
										}

									case 4:
									SearchLoop2:
										for {
											input := getMenu("Search by :", []string{"id", "title", "Category", "author", "View all"})
											switch input {
											case 0:
												break SearchLoop2
											case 1:
												id := getIntInput("Enter id")
												if id == 0 {
													break SearchLoop2
												}
												printBooksByID(id)
											case 2:
												title := getInputs([]string{"Enter title"})
												if title[0] == "0" {
													break SearchLoop2
												}
												printBooksByTitle(title[0])
											case 3:
												category := getInputs([]string{"Enter category"})
												if category[0] == "0" {
													break SearchLoop2
												}
												printBooksByCategory(category[0])
											case 4:
												author := getInputs([]string{"Enter author"})
												if author[0] == "0" {
													break SearchLoop2
												}
												printBooksByAuthor(author[0])
											case 5:
												printBooks()
											}

										}
									}
								}
							case 3:

							case 4:

							case 5:
								userInfo[0] = ""
								userInfo[1] = ""
								id = 0
								userid = ""
								break LoginLoop
							}
						}
					case "Author":
					AuthorMainLoop:
						for {
							fmt.Println("welcome Admin " + username)
							input := getMenu("Staff panel", []string{"Manage books", "View Sales"}) // set price , del add edit  print own book
							switch input {
							case 0:
								break AuthorMainLoop
							}
						}
					}
				}
			}
		case 2:
			inputs := getInputs([]string{"Enter username", "Enter email", "Enter password"})
			insertIntoUsers(append(inputs, "Customer"))
		}

	}
}
