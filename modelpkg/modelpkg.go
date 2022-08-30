package modelpkg

import (
	// "database/sql"
	"fmt"
	"strings"

	"github.com/aliaqa256/go_pgsql_orm/dbpkg"

	"reflect"
)

func Migrate(m any)   {
	database, err := dbpkg.NewDbConnection(
		"localhost",
		5432,
		"postgres",
		"",
		"test1",
	).CheckConnections()
	must(err)

	// //////////////
	model:= m

	tableName := reflect.TypeOf(model).Name()
	fmt.Println(tableName)
	filedsDict := make(map[string]string)
	filedsDict["id"] = "serial PRIMARY KEY"
	
	fileds := reflect.TypeOf(model)
	fmt.Println(fileds)
	for i := 0; i < fileds.NumField() ; i++ {
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

	sqlCmd:=fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %vs (`, strings.ToLower(tableName))
	counts:=0
	for key, value := range filedsDict {
		// if its last iteration
		if counts == len(filedsDict)-1 {
			sqlCmd += fmt.Sprintf(`%v %v)`, key, value)

		}else{
		sqlCmd += fmt.Sprintf("%s %s, ", key, value)}
		counts++
	}
	
	fmt.Println(sqlCmd)




		
	aaaa:=database.ExecuteCommand(sqlCmd)
	fmt.Println(aaaa.Resualt)
}

func must(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
