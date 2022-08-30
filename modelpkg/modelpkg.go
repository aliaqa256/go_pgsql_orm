package modelpkg

import (
	"fmt"
	"github.com/aliaqa256/go_pgsql_orm/dbpkg"
	"reflect"
	"strings"
)

func Migrate(m any) {
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
	filedsDict := make(map[string]string)
	filedsDict["id"] = "serial PRIMARY KEY"
	fileds := reflect.TypeOf(model)
	fmt.Println(fileds)
	for i := 0; i < fileds.NumField(); i++ {
		varName := fileds.Field(i).Name
		if varName == "Id" {
			continue
		}
		varType := fileds.Field(i).Type
		varTag := fileds.Field(i).Tag
		if varTag.Get("orm") == "" {
			switch varType.Name() {
			case "string":
				filedsDict[varName] = "varchar(255) NOT NULL"
			case "int":
				filedsDict[varName] = "integer"
			case "float":
				filedsDict[varName] = "float"
			case "Time":
				filedsDict[varName] = "timestamp with time zone NOT NULL DEFAULT now()"
			}
		} else {
			filedsDict[varName] = varTag.Get("orm")
		}

	}
	TabledoesExist := dbpkg.IsTableExist(database, strings.ToLower(tableName)+"s")
	if !TabledoesExist {
		sqlCmd := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %vs (`, strings.ToLower(tableName))
		counts := 0
		for key, value := range filedsDict {
			// if its last iteration
			if counts == len(filedsDict)-1 {
				sqlCmd += fmt.Sprintf(`%v %v)`, strings.ToLower(key), value)

			} else {
				sqlCmd += fmt.Sprintf("%v %v, ", strings.ToLower(key), value)
			}
			counts++
		}
		fmt.Println(sqlCmd)
		aaaa := database.ExecuteCommand(sqlCmd)
		fmt.Println(aaaa.Resualt)
	} else {
		for key, value := range filedsDict {
			database.ExecuteCommandAndIgnoreErrors(fmt.Sprintf(`ALTER TABLE %vs ADD COLUMN %v %v`, strings.ToLower(tableName), key, value))
			database.ExecuteCommandAndIgnoreErrors(fmt.Sprintf(`ALTER TABLE %vs ALTER COLUMN %v TYPE %v;`, strings.ToLower(tableName), key, strings.Split(value," ")[0] ))
		}
	}
}

func must(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
