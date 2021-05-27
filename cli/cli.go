package cli

import (
	"fmt"
	"github.com/tanakrid/pos/user"
	"github.com/tanakrid/pos/mock_data"
	"github.com/tanakrid/pos/util"
	"github.com/tanakrid/pos/stock"
	"errors"
	"log"
)

var employee user.Employee 

func headInterface(u user.User){
	var username string
	var password string
	fmt.Println("Hello, This is POS system for service to our customer to manage some transection.")
	fmt.Printf("Enter your username: ")
	fmt.Scanln(&username)
	fmt.Printf("Enter your password: ")
	fmt.Scanln(&password)

	if ! u.Authen(username, password, u) {
		log.Fatalln(errors.New("password incorrect"))
	}
	employee := user.Employee{User: user.User{Username: username, Password: password}}
	fmt.Printf("Have a nice day %s\n\n", employee.User.Username)
}

func bodyInterface(customer user.Customer, employee user.Employee){
	var operation string
	mock_stock := mock_data.GetStockData()

	for ;; {
		fmt.Printf("Enter operation: ")
		fmt.Scanln(&operation)
		if operation == "s" {
			return
		}
		if operation == "show" {
			fmt.Printf("%#v\n", mock_stock)

			fmt.Println(mock_stock.Name)

		}
		if operation == "cal" {
			var idProduct string
			fmt.Printf("Customer wallet: %v baht\nEnter product ID: ", customer.Wallet)
			fmt.Scanln(&idProduct)

			products := []stock.Product{}
			products = append(products, mock_stock.GetProduct(idProduct))

			transactionStatus, change := util.Pay(employee, customer, products, mock_stock)

			fmt.Printf("\nCustomer wallet : %v baht\n", customer.Wallet)
			if transactionStatus {
				fmt.Printf("\nCustomer change : %v baht\n", change)
			}
			
			fmt.Printf("Stock Data:\n%#v\n", mock_stock)
			
		}
	}
}

func ShowInterface(){
	mock_employee := mock_data.GetEmployeeData()
	customer := mock_data.GetCustomerData()
	headInterface(mock_employee.User)
	bodyInterface(customer, mock_employee)

}