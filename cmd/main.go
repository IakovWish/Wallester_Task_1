package main

import (
	"net/http"

	"github.com/IakovWish/Wallester_Task_1/customers"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/customers", customers.Index)
	http.HandleFunc("/customers/search", customers.Search)
	http.HandleFunc("/customers/order", customers.Order)
	http.HandleFunc("/customers/show", customers.Show)
	http.HandleFunc("/customers/create", customers.Create)
	http.HandleFunc("/customers/create/process", customers.CreateProcess)
	http.HandleFunc("/customers/edit", customers.Edit)
	http.HandleFunc("/customers/edit/process", customers.EditProcess)
	http.HandleFunc("/customers/delete/process", customers.DeleteProcess)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/customers", http.StatusSeeOther)
}
