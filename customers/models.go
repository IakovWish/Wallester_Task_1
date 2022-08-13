package customers

import (
	"errors"
	"net/http"

	"strconv"
	"time"

	"github.com/IakovWish/Wallester_Task_1/configs"
)

type Customer struct {
	Id         int
	First_name string    // required, max length 100
	Last_name  string    // required, max length 100
	Birth_date time.Time // required, from 18 till 60 years
	Gender     string    // required, allowed values are Male, Female
	E_mail     string    // required, should be valid email
	Address    string    // optional, max length 200
}

func AllCustomers() ([]Customer, error) {
	rows, err := configs.DB.Query("SELECT * FROM customers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	customers_arr := make([]Customer, 0)
	for rows.Next() {
		customer := Customer{}
		err := rows.Scan(
			&customer.Id,
			&customer.First_name,
			&customer.Last_name,
			//&customer.Birth_date,
			&customer.Gender,
			&customer.E_mail,
			&customer.Address)
		if err != nil {
			return nil, err
		}
		customers_arr = append(customers_arr, customer)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return customers_arr, nil
}

func OneCustomer(r *http.Request) (Customer, error) {
	customer := Customer{}
	id := r.FormValue("id")
	if id == "" {
		return customer, errors.New("400. Bad Request.")
	}

	row := configs.DB.QueryRow("SELECT * FROM customers WHERE id = $1", id)

	err := row.Scan(
		&customer.Id,
		&customer.First_name,
		&customer.Last_name,
		//&customer.Birth_date,
		&customer.Gender,
		&customer.E_mail,
		&customer.Address)

	if err != nil {
		return customer, err
	}

	return customer, nil
}

func PutCustomer(r *http.Request) (Customer, error) {
	// get form values
	customer := Customer{}
	customer.First_name = r.FormValue("first_name")
	customer.Last_name = r.FormValue("last_name")
	//customer.Birth_date = r.FormValue("birth_date")
	customer.Gender = r.FormValue("gender")
	customer.E_mail = r.FormValue("e_mail")
	customer.Address = r.FormValue("address")

	// validate form values
	if customer.First_name == "" || customer.Last_name == "" || customer.Gender == "" || customer.E_mail == "" /*customer.Birth_date == ""*/ {
		return customer, errors.New("400. Bad request. All fields must be complete.") // appart from Address
	}

	// insert values
	_, err := configs.DB.Exec("INSERT INTO customers (first_name, last_name, gender, e_mail, address) VALUES ($1, $2, $3, $4, $5)",
		customer.First_name, customer.Last_name /*customer.Birth_date,*/, customer.Gender, customer.E_mail, customer.Address)

	if err != nil {
		return customer, errors.New("500. Internal Server Error." + err.Error())
	}

	return customer, nil
}

func EditCustomer(r *http.Request) (Customer, error) {
	// get form values
	customer := Customer{}

	i := r.FormValue("id")
	customer.First_name = r.FormValue("first_name")
	customer.Last_name = r.FormValue("last_name")
	//customer.Birth_date = r.FormValue("birth_date")
	customer.Gender = r.FormValue("gender")
	customer.E_mail = r.FormValue("e_mail")
	customer.Address = r.FormValue("address")

	if customer.First_name == "" || customer.Last_name == "" || customer.Gender == "" || customer.E_mail == "" /*customer.Birth_date == ""*/ {
		return customer, errors.New("400. Bad request. All fields must be complete.") // appart from Address
	}

	// convert form values
	res, err := strconv.Atoi(i)
	if err != nil {
		panic(err)
	}
	customer.Id = res

	// insert values
	_, err = configs.DB.Exec("UPDATE customers SET first_name = $1, last_name = $2, gender = $3, e_mail = $4, address = $5 where id = $6",
		customer.First_name, customer.Last_name, customer.Gender, customer.E_mail, customer.Address, customer.Id)

	if err != nil {
		return customer, err
	}
	return customer, nil
}

func DeleteCustomer(r *http.Request) error {
	id := r.FormValue("id")
	if id == "" {
		return errors.New("400. Bad Request.")
	}

	_, err := configs.DB.Exec("DELETE FROM customers WHERE id=$1;", id)
	if err != nil {
		return errors.New("500. Internal Server Error")
	}
	return nil
}
