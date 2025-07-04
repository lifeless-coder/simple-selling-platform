package admin

import (
	"fmt"
	"simple_selling/db"
)

func showMsg() {
	fmt.Println("Hello Admin!")
	fmt.Println("1.Create new product")
	fmt.Println("2.Update existing product")
	fmt.Println("3.Delete existing product")
	fmt.Println("4.Show all existing product")
	fmt.Println("5.Logout")
	fmt.Println("Input your choice")

}
func takeInput() (int, string, float64) {
	fmt.Println("Input product ID")
	var input int
	fmt.Scan(&input)
	fmt.Println("Input product name")
	var productName string
	fmt.Scan(&productName)
	fmt.Println("Input product price")
	var productPrice float64
	fmt.Scan(&productPrice)
	return input, productName, productPrice
}
func AdminMain() {
	showMsg()
	var input int
	fmt.Scanf("%d", &input)

	switch input {
	case 1:
		id, name, price := takeInput()
		Create(id, name, price)
		AdminMain()
	case 2:
		id, name, price := takeInput()
		Update(id, name, price)
		AdminMain()
	case 3:
		id, _, _ := takeInput()
		Delete(id)
		AdminMain()
	case 4:
		db.Read()
		AdminMain()
	case 5:
		return
	default:
		fmt.Println("Invalid input")
		AdminMain()
	}
}
