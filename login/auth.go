package login

import "simple_selling/db"

func auth(name, pass string) bool {
	for n, p := range db.Info {
		if n == name && p == pass {
			return true
		}
	}
	return false
}
