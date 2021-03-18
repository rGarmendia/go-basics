package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name  string
	Email string
}

func allUsers(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("mysql", "root:my-secret-pw@tcp(localhost:3306)/testdb")
	if err != nil {
		fmt.Println(err.Error())
	}

	defer db.Close()

	var users []User
	db.Find(&users)
	fmt.Println("{}", users)

	json.NewEncoder(w).Encode(users)
}

func newUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "New User Endpoint Hit\n")

	db, err := gorm.Open("mysql", "root:my-secret-pw@tcp(localhost:3306)/testdb")

	if err != nil {
		panic("failed to connect database\n")
		fmt.Println(err.Error())
	}

	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]
	email := vars["email"]

	db.Create(&User{Name: name, Email: email})
	fmt.Fprintf(w, "New User Successfully Created")
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Delete User Endpoint Hit\n")

	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database\n")
	}
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]

	var user User

	db.Where("name = ?", name).Find(&user)
	db.Delete(&user)

	fmt.Fprintf(w, "Successfully Deleted User")
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Update User Endpoint Hit\n")

	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]
	email := vars["email"]

	var user User
	db.Where("name = ?", name).Find(&user)

	user.Email = email

	db.Save(&user)
	fmt.Fprintf(w, "Successfully Updated User")
}

func handleRequest() {
	myRouter := mux.NewRouter() //.StrictSlash(true)
	myRouter.HandleFunc("/users", allUsers).Methods("GET")
	myRouter.HandleFunc("/users/{name}", deleteUser).Methods("DELETE")
	myRouter.HandleFunc("/users/{name}/{email}", updateUser).Methods("PUT")
	myRouter.HandleFunc("/users/{name}/{email}", newUser).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", myRouter))

}

// initial migration function
func initialMigration() {
	db, err := gorm.Open("mysql", "root:my-secret-pw@tcp(localhost:3306)/testdb")
	if err != nil {
		fmt.Println(err.Error())
	}

	defer db.Close()

	db.AutoMigrate(&User{})
}

func main() {
	fmt.Println("GO ORM Tutorial")

	// Add the call to our new initialMigration func
	initialMigration()

	// Handle Subsequent requests
	handleRequest()
}
