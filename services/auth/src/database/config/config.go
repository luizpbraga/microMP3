package config

import (
	"errors"
	"fmt"
	"log"
	"os"
)

type mysqlENV string
type mapEnvs map[mysqlENV]string

func InitMySQLDataSorce() (string, error) {
	envs := mapEnvs{
		"MYSQL_PORT":     os.Getenv("MYSQL_PORT"),
		"MYSQL_USER":     os.Getenv("MYSQL_USER"),
		"MYSQL_HOST":     os.Getenv("MYSQL_HOST"),
		"MYSQL_DATABASE": os.Getenv("MYSQL_DATABASE"),
		"MYSQL_PASSWORD": os.Getenv("MYSQL_PASSWORD"),
	}

	log.Println(envs)

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
		HOST     = envs["MYSQL_HOST"]
		USER     = envs["MYSQL_USER"]
		DATABASE = envs["MYSQL_DATABASE"]
		PASSWORD = envs["MYSQL_PASSWORD"]
	)
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", USER, PASSWORD, HOST, DATABASE)
}
