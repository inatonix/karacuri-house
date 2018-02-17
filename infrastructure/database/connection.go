package infrastructure

import (
	"database/sql"
)

func ConnectDB() (db *sql.DB, err error) {
	db, err = sql.Open("mysql", "root:"+PASSWORD+"@unix(/cloudsql/"+PROJECT_NAME+":us-central1:karacuri-house/karacuri-house?charset=utf8&parseTime=True")
	return db, err
}
