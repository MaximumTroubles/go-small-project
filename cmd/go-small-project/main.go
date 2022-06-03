package main

import (
	"fmt"
	template2 "html/template"
	"net/http"
)

func main() {
	handleRequest()
}

type User struct {
	Name                 string
	Age                  uint8
	Balance              int
	AvgGrades, Happiness float64
	Hobbies              []string
}

func (u User) getAllInfo() string {
	return fmt.Sprintf("User name is: %s. He is %d old and he has %d money on balance",
		u.Name, u.Age, u.Balance)
}

func (u *User) setNewName(newName string) {
	u.Name = newName
}

func handleRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/contacts", contactsPage)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		return
	}
}

func contactsPage(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Contacts")
	if err != nil {
		return
	}
}

func homePage(w http.ResponseWriter, r *http.Request) {
	bob := User{
		Name:      "Bob",
		Age:       25,
		Balance:   0,
		AvgGrades: 4.2,
		Happiness: 0.8,
		Hobbies: []string{
			"football",
			"skate",
			"Dance",
		},
	}

	bob.setNewName("Max")
	templatePath, _ := template2.ParseFiles("../../web/templates/home_page.html")
	err := templatePath.Execute(w, bob)
	if err != nil {
		return
	}
}
