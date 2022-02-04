package util

import (
	"database/sql"
	"fmt"
	"sirclo/config"

	_ "github.com/go-sql-driver/mysql"
)

func MysqlDriver(config *config.AppConfig) *sql.DB {
	uri := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true",
		config.Database.Username,
		config.Database.Password,
		config.Database.Address,
		config.Database.Port,
		config.Database.Name)
	db, err := sql.Open("mysql", uri)
	if err != nil {
		err = fmt.Errorf("failed to connect database")
		panic(err)
	}
	return db
}
