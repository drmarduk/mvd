package Db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

var (
	c *sql.DB
)
type tblNoten struct{
	Id int
	Name string
	dir string
	Created time.Time
}

func init() {
	c, err := sql.Open("mysql", "root:mvd@/mvd")
	if err != nil {

		log.Printf("Db.init: Error while creating Connection: %s\n", err.Error())
		return
	}
	log.Printf("Try to create table");
	_, err = c.Exec("create table notensatz (id integer not null primary key auto_increment);")
	if err != nil {
		log.Printf("Db.init: %s\n", err.Error())
		return
	}

	log.Printf("Table created.\n")
}
func Hi() {
	log.Printf("hi")
}