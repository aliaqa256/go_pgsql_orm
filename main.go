package main

import (
	"fmt"
	managers "github.com/aliaqa256/go_pgsql_orm/managerspkg"
	"github.com/aliaqa256/go_pgsql_orm/modelpkg"
	"time"
)

type User struct {
	Id        uint `orm:"serial PRIMARY KEY"`
	CreatedAt time.Time
	UpdatedAt time.Time
	UserName  string  `orm:"varchar(255) not null"`
	Email     string  `orm:"varchar(255) not null"`
	Phone     string  `orm:"varchar(255) not null"`
	Age       float32 `orm:"int"`
}

func main() {
	a := User{
		UserName: "Aliaqa",
		Email:    "1",
		Phone:    "1",
		Age:      1,
	}
	modelpkg.Migrate(a)
	managers.Create(a)
	dbc := managers.GetId(a, map[string]string{
		"id": "2",
	})
	fmt.Println(dbc.Resualt)
	dbc = managers.GetField(a, 26, "age")
	fmt.Println(dbc.Resualt)

}
