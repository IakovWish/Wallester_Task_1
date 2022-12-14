package customers

import (
	"database/sql"
	"net/http"

	"github.com/IakovWish/Wallester_Task_1/configs"
)

func Index(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	customers_arr, err := AllCustomers(r)

	if err != nil {
		if err.Error() == "page is not available" {
			http.Redirect(w, r, "/customers?ord=id&page=1", http.StatusSeeOther)
		}

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	configs.TPL.ExecuteTemplate(w, "customers.gohtml", customers_arr)
}

func Search(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	customers_arr, err := SearchedCustomers(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	configs.TPL.ExecuteTemplate(w, "customers.gohtml", customers_arr)
}

func Show(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	customer, err := OneCustomer(r)
	switch {
	case err == sql.ErrNoRows:
		http.NotFound(w, r)
		return
	case err != nil:
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	configs.TPL.ExecuteTemplate(w, "show.gohtml", customer)
}

func Create(w http.ResponseWriter, r *http.Request) {
	configs.TPL.ExecuteTemplate(w, "create.gohtml", nil)
}

func CreateProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	_, err := PutCustomer(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotAcceptable)
		return
	}

	http.Redirect(w, r, "/customers?ord=id&page=1", http.StatusSeeOther)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	customer, err := OneCustomer(r)
	switch {
	case err == sql.ErrNoRows:
		http.NotFound(w, r)
		return
	case err != nil:
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	configs.TPL.ExecuteTemplate(w, "edit.gohtml", customer)
}

func EditProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	_, err := EditCustomer(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, "/customers?ord=id&page=1", http.StatusSeeOther)
}

func DeleteProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	err := DeleteCustomer(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, "/customers?ord=id&page=1", http.StatusSeeOther)
}
