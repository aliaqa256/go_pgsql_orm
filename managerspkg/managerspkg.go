package managers

import (
	"fmt"
	"reflect"
	"strings"
	"github.com/aliaqa256/go_pgsql_orm/dbpkg"
	"github.com/fatih/structs"
)

func  Create(m any){
	database, err := dbpkg.NewDbConnection(
		"localhost",
		5432,
		"postgres",
		"",
		"test1",
	).CheckConnections()
	must(err)

	model := m
	tableName := reflect.TypeOf(model).Name()
	fmt.Println(tableName)

	// struct to map
	newStruct := structs.New(m)
	MapOfModel:=newStruct.Map()

	// delete id from map
	delete(MapOfModel, "Id")
	delete(MapOfModel, "CreatedAt")
	delete(MapOfModel, "UpdatedAt")
	
	sliceOfkeys := make([]string, 0)
	sliceOfvalues := make([]string, 0)
	for key , value  := range MapOfModel {
		sliceOfkeys = append(sliceOfkeys,strings.ToLower( key))
		sliceOfvalues = append(sliceOfvalues, fmt.Sprintf("'%v'",value))
	}
	sqlCmd:=fmt.Sprintf("INSERT INTO %vs",strings.ToLower(tableName))
	sqlCmd += " ("
	// join keys with ,
	keyssss:=strings.Join(sliceOfkeys,`, `)
	valuessss:=strings.Join(sliceOfvalues,`, `)
	sqlCmd += keyssss
	sqlCmd += ") VALUES ("
	sqlCmd += valuessss
	sqlCmd += ");"
	fmt.Println(sqlCmd)
	// execute sql command
	database.ExecuteCommand(sqlCmd)
}

func must(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
