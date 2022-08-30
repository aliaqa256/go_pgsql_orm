package main

import (
	"fmt"
	"time"
	"github.com/aliaqa256/go_pgsql_orm/modelpkg"
	// "github.com/aliaqa256/go_pgsql_orm/dbpkg"
)





type User struct {

	Id uint  `orm:"serial PRIMARY KEY"`
	CreatedAt time.Time 
	UpdatedAt time.Time 
	Namea string `orm:"varchar(255) not null"`
	Email string `orm:"varchar(255) not null"`
}




func main() {
	// database, err :=dbpkg.NewDbConnection(
	// 	"localhost",
	// 	5432,
	// 	"postgres",
	// 	"",
	// 	"test1",
	// ).CheckConnections()
	// must(err)
	// database.ExecuteCommand("CREATE TABLE IF NOT EXISTS users (id serial PRIMARY KEY, created_at timestamp with time zone NOT NULL DEFAULT now(), updated_at timestamp with time zone NOT NULL DEFAULT now(), name varchar(255) NOT NULL, email varchar(255) NOT NULL)")
	
	a := User{}
	modelpkg.Migrate(a)

	






}


func must(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}