package NotenSatz

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"time"
)

type NotenSatz struct {
	Id      int
	Name    string
	Dir     string
	Created time.Time
}

func Setup() bool {
	c, err := sql.Open("sqlite3", "mvd.db")
	if err != nil {
		log.Printf("NotenSatz.Setup: %s\n", err.Error())
		return false
	}
	defer c.Close()
	query := "create table if not exists notensatz(id integer not null primary key autoincrement" +
		", name text, dir text, created timestamp default current_timestamp);"
	_, err = c.Exec(query)
	if err != nil {
		log.Printf("NotenSatz.Setup: %s\n", err.Error())
		return false
	}
	return true
}
func New(name string) bool {
	c, err := sql.Open("sqlite3", "mvd.db")
	if err != nil {
		log.Printf("NotenSatz.New: %s\n", err.Error())
		return false
	}
	defer c.Close()
	query := "insert into notensatz(name) values(?)"
	_, err = c.Exec(query, name)
	if err != nil {
		log.Printf("NotenSatz.New: %s\n", err.Error())
		return false
	}
	return true
}

func Get(id int) (n NotenSatz, err error) {
	c, err := sql.Open("sqlite3", "mvd.db")
	if err != nil {
		log.Printf("NotenSatz.Get: %s\n", err.Error())
		return n, err
	}
	defer c.Close()
	query := "select * from notensatz where id = ?"
	row := c.QueryRow(query, id)
	err = row.Scan(&n.Id, &n.Name, &n.Dir, &n.Created)
	if err != nil {
		log.Printf("NotenSatz.Get: %s\n", err.Error())
		return n, err
	}
	return n, nil
}
