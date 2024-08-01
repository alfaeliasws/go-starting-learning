package main

import (
	"fmt"
	go_mysql_api "go-mysql"
)

func main() {
	go_mysql_api.Create()
	go_mysql_api.GetAll()
	go_mysql_api.Update(1, "Joni")
	go_mysql_api.Delete(4)
	go_mysql_api.GetById(2)

	fmt.Println("Success")
}
