package config

import (
	"errors"
	"fmt"
	"os"
)

type mysqlENV string
type mapEnvs map[mysqlENV]string

func InitMySQLDataSorce() (string, error) {
	envs := mapEnvs{
		"MYSQL_DB":       os.Getenv("MYSQL_DB"),
		"MYSQL_PORT":     os.Getenv("MYSQL_PORT"),
		"MYSQL_USER":     os.Getenv("MYSQL_USER"),
		"MYSQL_HOST":     os.Getenv("MYSQL_HOST"),
		"MYSQL_PASSWORD": os.Getenv("MYSQL_PASSWORD"),
	}

	if err := envs.check(); err != nil {
		return "", err
	}

	return envs.format(), nil
}

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
