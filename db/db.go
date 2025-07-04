package db

import "fmt"

//import (
//	"fmt"
//)

//type Person struct {
//	Name string
//	Age  int
//	pass string
//}
//
//func info() {
//	var Admin Person
//	var User Person
//
//	Admin.Age = 10
//	Admin.pass = "123456"
//	Admin.Name = "Adrita"
//
//	User.Age = 10
//	User.pass = "123456"
//	User.Name = "Ramim"
//}
//func (p *Person) getInfo() (string, string) {
//	info()
//	return p.Name, p.pass
//}

var Info = map[string]string{
	"Adrita": "123456",
	"Ramim":  "123456",
}

type Product struct {
	Name  string
	Price float64
}

var Item = map[int]Product{
	1: Product{Name: "shoe", Price: 2000},
	2: Product{Name: "car", Price: 200000},
}
var UserBal float64 = 10000

func Read() {
	for key, product := range Item {
		fmt.Println("Product key: ", key, "product name : ", product.Name, "product price : ", product.Price)
	}
}
