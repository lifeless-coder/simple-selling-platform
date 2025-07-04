package user

import (
	"fmt"
	"simple_selling/db"
)

func showMsg() {
	fmt.Println("hello user!")
	fmt.Println("1.Buy Products")
	fmt.Println("2.show balance")
	fmt.Println("3.Add Balance")
	fmt.Println("4.logout")
	fmt.Println("Input your choice")

}

func takeOrder(bal float64) float64 {
	db.Read()
	fmt.Println("Input Product Id to buy: ")
	var productId int
	fmt.Scan(&productId)
	value, exist := db.Item[productId]

	if exist {

		if value.Price > bal {
			fmt.Println("Incufficient balance")
			return 0
		} else {
			fmt.Println("Your order done. ")
			return value.Price
		}

	} else {
		fmt.Println("Invalid ID")
		return 0
	}

	return 0
}
