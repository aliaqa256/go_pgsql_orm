package main

import (
	"fmt"
	"time"
	"github.com/aliaqa256/go_pgsql_orm/modelpkg"
)





type User struct {

	Id uint  `orm:"serial PRIMARY KEY"`
	CreatedAt time.Time 
	UpdatedAt time.Time 
	Name string `orm:"varchar(255) not null"`
	Email string `orm:"varchar(255) not null"`
	Phone string `orm:"varchar(255) not null"`
	Age  float32  `orm:"int"`
}




func main() {
	a := User{}
	modelpkg.Migrate(a)
}


func must(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}