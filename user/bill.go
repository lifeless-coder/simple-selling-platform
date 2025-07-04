package user

import (
	"fmt"
	"simple_selling/db"
)

func addBalance(prev float64) float64 {
	fmt.Println("Add Balance: ")
	var input float64
	fmt.Scanf("%f", &input)
	total := prev + input
	fmt.Println("Your current balance is: ", total)
	return total
}
func calculatebal(prev float64, price float64) float64 {
	balance := prev - price
	db.UserBal = balance
	return balance
}
