package util

import (
	"github.com/tanakrid/pos/user"
	"github.com/tanakrid/pos/stock"
)

func Pay(employee user.Employee, customer user.Customer, products []stock.Product, stock stock.Stock) (bool, int) {
	cash := customer.Wallet
	var amount int
	for i := 0; i < len(products); i++ {
		amount += products[i].Price
	}
	transactionComplete := cash >= amount
	if (!transactionComplete) {
		return transactionComplete, 0
	}
	cash -= amount
	customer.Wallet = cash
	for i := 0; i < len(products); i++ {
		stock.RemoveProduct(products[i].IdProduct)
	}
	return transactionComplete, cash
}