package mock_data

import "github.com/tanakrid/pos/user"
import "github.com/tanakrid/pos/stock"

func GetStockData() stock.Stock {
	products := []stock.Product{}
	products = append(products, stock.Product{Name: "snack", IdProduct: "12345", Price: 200, TypeProduct: "food"})
	products = append(products, stock.Product{Name: "salad", IdProduct: "12346", Price: 300, TypeProduct: "food"})
	products = append(products, stock.Product{Name: "water", IdProduct: "12347", Price: 100, TypeProduct: "drink"})

	stock := stock.Stock{Name: "stock1", Products: products}
	return stock
}

func GetEmployeeData() user.Employee {
	return user.Employee{User: user.User{Username: "sunny", Password: "111111"}}
}

func GetCustomerData() user.Customer {
	return user.Customer{User: user.User{Username: "tanakrid", Password: "123456"}, Wallet: 5000}
}