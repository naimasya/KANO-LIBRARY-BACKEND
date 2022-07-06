package main

import (
	"fmt"
	"kanolibrary/router"
	"kanolibrary/util"
	"net/http"
)

func main() {
	// crud.Insert("books", books{Name: "Hujan", Author: "Tere Liye"})
	// crud.Find("books", books{})
	// crud.Update()
	http.HandleFunc("/books", router.FindAll)
	http.HandleFunc("/books/find", router.FindByID)
	http.HandleFunc("/books/insert", router.InsertBooks)
	http.HandleFunc("/books/update", router.UpdateBook)
	http.HandleFunc("/books/delete", router.DeleteBook)
	fmt.Println("starting web server at http://localhost:8080/")
	err := http.ListenAndServe(":8080", nil)
	util.ErrorChecker(err)
}
