package dbpkg

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type DbConnection struct {
	Host     string
	Port     int
	User     string
	Password string
	Dbname   string
	ConnectionString string
	Resualt any
}

func (db *DbConnection) QuaryCommandSelect(cmd string) *DbConnection {
	database, err := sql.Open("postgres", db.ConnectionString)
	must(err)
	defer database.Close()
	// QueryRow 
	var res any
	err = database.QueryRow(cmd).Scan(&res)
	must(err)
	db.Resualt = res
	return db
}

func (db *DbConnection) ExecuteCommandAndIgnoreErrors(cmd string) *DbConnection {
	database, err := sql.Open("postgres", db.ConnectionString)
	must(err)
	defer database.Close()
	_, err = database.Exec(cmd)
	ignor(err)

	return db
}

// note cheked this func yet
func (db *DbConnection) ExecuteQueryRows(cmd string) *DbConnection {
	database, err := sql.Open("postgres", db.ConnectionString)
	must(err)
	defer database.Close()
	var res any
	err = database.QueryRow(cmd).Scan(&res)
	must(err)
	db.Resualt = res
	return db

}

func IsTableExist(db *DbConnection, tableName string) bool {
	database, err := sql.Open("postgres", db.ConnectionString)
	must(err)
	defer database.Close()
	var count int
	err = database.QueryRow("SELECT COUNT(*) FROM information_schema.tables WHERE table_name = $1", tableName).Scan(&count)
	if err != nil {
		return false
	}
	return count > 0
}

func (db *DbConnection) ExecuteCommand(cmd string) *DbConnection {
	database, err := sql.Open("postgres", db.ConnectionString)
	must(err)
	defer database.Close()
	_, err = database.Exec(cmd)
	must(err)

	return db
}

// note usefull yet
func (db *DbConnection) QuaryCommand(cmd string) *DbConnection {
	database, err := sql.Open("postgres", db.ConnectionString)
	must(err)
	defer database.Close()
	res, err := database.Query(cmd)
	must(err)
	db.Resualt = res
	return db
}

func (db *DbConnection) CheckConnections() (*DbConnection, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		db.Host, db.Port, db.User, db.Password, db.Dbname)
	db.ConnectionString = psqlInfo
	db.Resualt = nil
	database, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	defer database.Close()
	err = database.Ping()
	if err != nil {
		return nil, err

	}
	fmt.Println("db connected")
	return db, nil
}

// singelton
var initinatedConnection *DbConnection = nil
func NewDbConnection(host string, port int, user string, password string, dbname string) *DbConnection {
	if initinatedConnection != nil {
		return initinatedConnection
	}
	initinatedConnection = &DbConnection{
		Host:     host,
		Port:     port,
		User:     user,
		Password: password,
		Dbname:   dbname,
	}

	return initinatedConnection
}



func must(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func ignor(err error) {
	if err != nil {
		fmt.Println(err)
	}
}