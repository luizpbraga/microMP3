package database

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

type mysqlENV string

type mapEnvs map[mysqlENV]string

func (envs mapEnvs) check() error {
	msg := ""

	for key, val := range envs {
		if val == "" {
			msg += fmt.Sprintf("Missing Environment Variable: %s\n", key)
		}
	}

	if msg != "" {
		return errors.New(msg)
	}

	return nil
}

func (envs mapEnvs) format() string {
	var (
		DB       = envs["MYSQL_DB"]
		PORT     = envs["MYSQL_PORT"]
		USER     = envs["MYSQL_USER"]
		HOST     = envs["MYSQL_HOST"]
		PASSWORD = envs["MYSQL_PASSWORD"]
	)

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", USER, PASSWORD, HOST, PORT, DB)
}

func initMySQLDataSorce() (string, error) {
	envs := mapEnvs{
		"MYSQL_DB":       os.Getenv("MYSQL_DB"),
		"MYSQL_PORT":     os.Getenv("MYSQL_PORT"),
		"MYSQL_USER":     os.Getenv("MYSQL_USER"),
		"MYSQL_HOST":     os.Getenv("MYSQL_HOST"),
		"MYSQL_PASSWORD": os.Getenv("MYSQL_PASSWORD"),
	}

	err := envs.check()

	if err != nil {
		return "", err
	}

	return envs.format(), nil

}

var Db *sql.DB

func InitDataBase() (*sql.DB, error) {
	connection_string, err := initMySQLDataSorce()

	if err != nil {
		return nil, err
	}

	fmt.Print(connection_string)

	Db, err = sql.Open("mysql", connection_string)

	return Db, err
}
