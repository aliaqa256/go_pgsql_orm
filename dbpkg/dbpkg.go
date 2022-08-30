package dbpkg

import (
	"fmt"
	"database/sql"
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

func (db *DbConnection) ExecuteCommand(cmd string) *DbConnection {
	database, err := sql.Open("postgres", db.ConnectionString)
	must(err)
	defer database.Close()
	res, err := database.Exec(cmd)
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

func NewDbConnection(host string, port int, user string, password string, dbname string) *DbConnection {
	return &DbConnection{
		Host:     host,
		Port:     port,
		User:     user,
		Password: password,
		Dbname:   dbname,
	}
}



func must(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}