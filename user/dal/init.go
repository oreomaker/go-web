package dal

import "go-web/user/dal/db"

// Init init dal
func Init() {
	db.Init() // mysql init
}
