package user

import (
	"fmt"
	"simple_selling/db"
)

func UserMain() {
	bal := db.UserBal
	showMsg()
	var choice int
	fmt.Scanf("%d", &choice)
	fmt.Println(choice)
	switch choice {
	case 1:
		db.Read()
		var price float64
		price = takeOrder(bal)
		calculatebal(bal, price)
		bal = db.UserBal

		UserMain()
	case 2:
		fmt.Println("Balance: ", bal)
		UserMain()
	case 3:
		bal = addBalance(bal)
		bal = db.UserBal
		UserMain()

	case 4:
		return
	default:
		fmt.Println("choice not recognized")
		UserMain()
	}
}
