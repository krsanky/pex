package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/pelletier/go-toml"
	"go.d34d.net/pex/lg"
)

var DB *sql.DB

func Init(settings *toml.Tree) {
	lg.Log.Printf("init pg db start ...")

	password := settings.Get("db.password").(string)
	user := settings.Get("db.user").(string)
	db_name := settings.Get("db.name").(string)

	var err error
	connect_string := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		user, password, db_name)

	DB, err = sql.Open("postgres", connect_string)
	if err != nil {
		panic(err)
	}
	if DB.Ping() != nil {
		panic(err)
	}
}

func Drivers() {
	for _, d := range sql.Drivers() {
		fmt.Printf("driver:%s\n", d)
	}
}

func TestDB() {
	var err error
	if err = DB.Ping(); err != nil {
		panic(err)
	}

	rows, err := DB.Query("SELECT 'test'")
	defer rows.Close()
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			panic(err)
		}
		lg.Log.Printf("TestDB() name:%s\n", name)
	}
	if err := rows.Err(); err != nil {
		panic(err)
	}
}
