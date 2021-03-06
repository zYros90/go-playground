package main

import (
	"fmt"
	"log"
	"requests/examples"
)

func main() {
	employees, err := examples.GetExampleEmployees("https://dummy.restapigithub.com/zYros90/go-playground/api/v1/employees")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("number of employees: ", len(employees.Data))

	employee, err := examples.PostExampleEmployees("https://dummy.restapigithub.com/zYros90/go-playground/api/v1/create")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(employee)
}
