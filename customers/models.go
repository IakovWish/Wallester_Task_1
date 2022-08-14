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

func AllCustomers(r *http.Request) ([]Customer, error) {
	//var query string
	//fmt.Println(r.FormValue("srch_first"))
	//fmt.Println(r.FormValue("srch_last"))
	// if r.FormValue("srch_first") == "" && r.FormValue("srch_last") == "" {
	// 	fmt.Println("1")
	// 	query = "SELECT * FROM customers"
	// } else if r.FormValue("srch_first") != "" && r.FormValue("srch_last") != "" {
	// 	fmt.Println("2")
	// 	query = "SELECT * FROM customers WHERE first_name = " + r.FormValue("srch_first") + "AND last_name = " + r.FormValue("srch_last") + ";"
	// } else {
	// 	fmt.Println("3")
	// }

	//fmt.Println(query)

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
			&customer.Birth_date,
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

func SearchedCustomers(r *http.Request) ([]Customer, error) {

	customers_arr := make([]Customer, 0)
	if r.FormValue("srch_first") == "" || r.FormValue("srch_last") == "" {
		return customers_arr, errors.New("400. Bad Request")
	}

	rows, err := configs.DB.Query("SELECT * FROM customers WHERE first_name = $1 AND last_name = $2;",
		r.FormValue("srch_first"), r.FormValue("srch_last"))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	customers_arr = make([]Customer, 0)
	for rows.Next() {
		customer := Customer{}
		err := rows.Scan(
			&customer.Id,
			&customer.First_name,
			&customer.Last_name,
			&customer.Birth_date,
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
		return customer, errors.New("400. Bad Request")
	}

	row := configs.DB.QueryRow("SELECT * FROM customers WHERE id = $1", id)

	err := row.Scan(
		&customer.Id,
		&customer.First_name,
		&customer.Last_name,
		&customer.Birth_date,
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
	convert_date := r.FormValue("birth_date")
	customer.Gender = r.FormValue("gender")
	customer.E_mail = r.FormValue("e_mail")
	customer.Address = r.FormValue("address")

	// validate form values
	if customer.First_name == "" || customer.Last_name == "" || customer.Gender == "" || customer.E_mail == "" || convert_date == "" {
		return customer, errors.New("400. Bad request. All fields must be complete") // appart from Address
	}

	dateString := "2006-01-02"
	date_res, err := time.Parse(dateString, convert_date)
	if err != nil {
		panic(err)
	}
	customer.Birth_date = date_res

	if age(customer.Birth_date, time.Now()) < 18 || age(customer.Birth_date, time.Now()) > 60 {
		return customer, errors.New("400. Bad request. Age must befrom 18 till 60 years")
	}

	// insert values
	_, err = configs.DB.Exec("INSERT INTO customers (first_name, last_name, birth_date, gender, e_mail, address) VALUES ($1, $2, $3, $4, $5, $6)",
		customer.First_name, customer.Last_name, customer.Birth_date, customer.Gender, customer.E_mail, customer.Address)

	if err != nil {
		return customer, errors.New("500. Internal Server Error." + err.Error())
	}

	return customer, nil
}

func EditCustomer(r *http.Request) (Customer, error) {
	// get form values
	customer := Customer{}
	convert_id := r.FormValue("id")
	customer.First_name = r.FormValue("first_name")
	customer.Last_name = r.FormValue("last_name")
	convert_date := r.FormValue("birth_date")
	customer.Gender = r.FormValue("gender")
	customer.E_mail = r.FormValue("e_mail")
	customer.Address = r.FormValue("address")

	if customer.First_name == "" || customer.Last_name == "" || customer.Gender == "" || customer.E_mail == "" || convert_date == "" {
		return customer, errors.New("400. Bad request. All fields must be complete") // appart from Address
	}

	// convert form values
	id_res, err := strconv.Atoi(convert_id)
	if err != nil {
		panic(err)
	}
	customer.Id = id_res

	dateString := "2006-01-02"
	date_res, err := time.Parse(dateString, convert_date)
	if err != nil {
		panic(err)
	}
	customer.Birth_date = date_res

	if age(customer.Birth_date, time.Now()) < 18 || age(customer.Birth_date, time.Now()) > 60 {
		return customer, errors.New("400. Bad request. Age must befrom 18 till 60 years")
	}

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
		return errors.New("400. Bad Request")
	}

	_, err := configs.DB.Exec("DELETE FROM customers WHERE id=$1;", id)
	if err != nil {
		return errors.New("500. Internal Server Error")
	}
	return nil
}

func age(birthdate, today time.Time) int {
	today = today.In(birthdate.Location())
	ty, tm, td := today.Date()
	today = time.Date(ty, tm, td, 0, 0, 0, 0, time.UTC)
	by, bm, bd := birthdate.Date()
	birthdate = time.Date(by, bm, bd, 0, 0, 0, 0, time.UTC)
	if today.Before(birthdate) {
		return 0
	}
	age := ty - by
	anniversary := birthdate.AddDate(age, 0, 0)
	if anniversary.After(today) {
		age--
	}
	return age
}
