package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type DataBase struct {
	Host     string
	Port     int
	Database string
	User     string
	Password string
	Conn     *sql.DB
}

func (db *DataBase) Connect() {
	sqlDsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", db.User, db.Password, db.Host, db.Port, db.Database)
	conn, err := sql.Open("mysql", sqlDsn)
	if err != nil {
		dbConn.NotResultQueryExec(fmt.Sprintf("INSERT INTO E_LOG(MAC_ID, LOG, CREATE_DATE) VALUES ('system', '%s', NOW());", err))
		panic(err)
	}
	db.Conn = conn
}

func (db *DataBase) Close() {
	db.Conn.Close()
	fmt.Println("\r=> Close Database <=")
}

func (db *DataBase) NotResultQueryExec(sql string) bool {
	_, err := db.Conn.Exec(sql)
	if err != nil {
		fmt.Println(err)
		return false
	}

	return true
}
