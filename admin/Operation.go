package admin

import (
	"fmt"
	"simple_selling/db"
)

func Create(id int, name string, price float64) {
	value, exist := db.Item[id]
	_ = value
	if exist {
		fmt.Println("Item already exists")

	} else {
		db.Item[id] = db.Product{Name: name, Price: price}
	}
}
func Update(id int, name string, price float64) {
	value, exist := db.Item[id]
	_ = value
	if exist {
		db.Item[id] = db.Product{Name: name, Price: price}

	} else {
		fmt.Println("Item is not exists")
	}
}
func Delete(id int) {
	value, exist := db.Item[id]
	_ = value
	if exist {
		delete(db.Item, id)
	} else {
		fmt.Println("Item is not exists")
	}
}
