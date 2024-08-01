package go_mysql_api

import (
	"context"
	"fmt"
	"strconv"
)

func Create() {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "INSERT INTO customers(name) VALUES('Eko')"

	_, err := db.ExecContext(ctx, query)

	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new customer")

}

func GetAll() {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "SELECT id, name FROM customers"

	rows, err := db.QueryContext(ctx, query)

	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new customer")

	defer rows.Close()

	for rows.Next() {
		var (
			id   int16
			name string
		)
		err = rows.Scan(&id, &name)

		if err != nil {
			panic(err)
		}

		fmt.Println("Id:", id)
		fmt.Println("Name:", name)

	}
}

func GetById(id int16) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "SELECT id, name FROM customers WHERE id = " + strconv.Itoa(int(id))

	rows, err := db.QueryContext(ctx, query)

	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new customer")

	defer rows.Close()

}

func Update(id int16, arg string) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "UPDATE customers SET updated_at = NOW(), name = '" + arg + "'  WHERE id = " + strconv.Itoa(int(id))

	rows, err := db.QueryContext(ctx, query)

	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new customer")

	defer rows.Close()

}

func Delete(id int16) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "UPDATE customers SET deleted_at = NOW()  WHERE id = " + strconv.Itoa(int(id))

	rows, err := db.QueryContext(ctx, query)

	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new customer")

	defer rows.Close()

}
