package nyx
// This package handles all database operations for every service
// like the router, tasks, sub-tasks, etc
// This package is a wrapper using GORM
// https://github.com/jinzhu/gorm
// But why is it called Nyx? Well the connection of Nyx and Erebos
// brought forth Hemera (the DAY) and Aither (the AIR)


import (
	"github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

type Connection struct {
	User string
    Pw string
    Db string
}

func (t *Connection) Open() (*gorm.DB, error) {
    var err error
    var conn = t.User + ":" + t.Pw + "@/" + t.Db + "?charset=utf8&parseTime=True&loc=Local"

	db, err := gorm.Open("mysql", conn)

    return db, err
}
