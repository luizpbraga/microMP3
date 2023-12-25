package connection

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/luizpbraga/microMP3/services/auth/src/database/config"
)

var Db *sql.DB

func InitDataBase() (*sql.DB, error) {
	connection_string, err := config.InitMySQLDataSorce()
	if err != nil {
		return nil, err
	}

	fmt.Print(connection_string)

	Db, err = sql.Open("mysql", connection_string)

	return Db, err
}
