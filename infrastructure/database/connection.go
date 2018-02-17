package infrastructure

import "database/sql"

func ConnectDB() (db *sql.DB, err error) {
	db, err = sql.Open("mysql", "root:@unix(/cloudsql/karacuri-house:us-central1:karacuri-house)/karacuri-house?charset=utf8&parseTime=True")
	return db, err
}
