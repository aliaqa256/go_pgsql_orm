package main

import (
	"fmt"
	"time"
	// "github.com/aliaqa256/go_pgsql_orm/modelpkg"
	managers "github.com/aliaqa256/go_pgsql_orm/managerspkg"
)





type bbb struct {

	Id uint  `orm:"serial PRIMARY KEY"`
	CreatedAt time.Time 
	UpdatedAt time.Time 
	UserName string `orm:"varchar(255) not null"`
	Email string `orm:"varchar(255) not null"`
	Phone string `orm:"varchar(255) not null"`
	Age  float32  `orm:"int"`
}




func main() {
	a := bbb{
		UserName: "Aliaqa",
		Email: "1",
		Phone: "1",
		Age: 1,
	}
	// modelpkg.Migrate(a)
	managers.Create(a)





}


func must(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}