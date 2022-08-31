package managers

import (
	"fmt"
	"github.com/aliaqa256/go_pgsql_orm/dbpkg"
	"github.com/fatih/structs"
	"reflect"
	"strings"
)

func Create(m any) {
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
	MapOfModel := newStruct.Map()
	// delete id from map
	delete(MapOfModel, "Id")
	delete(MapOfModel, "CreatedAt")
	delete(MapOfModel, "UpdatedAt")

	sliceOfkeys := make([]string, 0)
	sliceOfvalues := make([]string, 0)
	for key, value := range MapOfModel {
		sliceOfkeys = append(sliceOfkeys, strings.ToLower(key))
		sliceOfvalues = append(sliceOfvalues, fmt.Sprintf("'%v'", value))
	}
	sqlCmd := fmt.Sprintf("INSERT INTO %vs", strings.ToLower(tableName))
	sqlCmd += " ("
	// join keys with ,
	keyssss := strings.Join(sliceOfkeys, `, `)
	valuessss := strings.Join(sliceOfvalues, `, `)
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

func GetId(m any, args map[string]string) *dbpkg.DbConnection {
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

	sqlCmd := fmt.Sprintf("SELECT id FROM %vs", strings.ToLower(tableName))
	sqlCmd += " WHERE "
	count := 0
	for key, value := range args {
		if count == len(args)-1 {
			sqlCmd += fmt.Sprintf("%v = '%v'", key, value)
		} else {
			sqlCmd += fmt.Sprintf("%v = '%v'", key, value)
			sqlCmd += " and "
		}
		count++
	}
	sqlCmd += ";"
	dbc := database.QuaryCommandSelect(sqlCmd)
	return dbc
}

func GetField(m any, id int, arg string) *dbpkg.DbConnection {
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
	sqlCmd := " SELECT "
	sqlCmd += fmt.Sprintf("%v", arg)
	sqlCmd += fmt.Sprintf(" FROM %vs ", strings.ToLower(tableName))
	sqlCmd += fmt.Sprintf(" WHERE id = %v", id)
	sqlCmd += ";"
	dbc := database.QuaryCommandSelect(sqlCmd)
	return dbc
}
