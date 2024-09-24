package models

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	_ "github.com/lib/pq"
)

type Config struct {
	DB *sql.DB
	DBname string
	DBpassword string
	DBuser string
	DBhost string
	DBport string
	DBconn string
}

var NewConfig Config

func InitConfig() error {
	var(
		err error
	)
	NewConfig.DBname = os.Getenv("DB_NAME")
	if NewConfig.DBname == "" {
		err = errors.New("DB_NAME is not set")
		return err
	}
	NewConfig.DBpassword = os.Getenv("DB_PASSWORD")
	if NewConfig.DBpassword == "" {
		err = errors.New("DB_PASSWORD is not set")
		return err
	}
	NewConfig.DBuser = os.Getenv("DB_USER")
	if NewConfig.DBuser == "" {
		err = errors.New("DB_USER is not set")
		return err
	}
	NewConfig.DBhost = os.Getenv("DB_HOST")
	if NewConfig.DBhost == "" {
		err = errors.New("DB_HOST is not set")
		return err
	}
	NewConfig.DBport = os.Getenv("DB_PORT")
	if NewConfig.DBport == "" {
		err = errors.New("DB_PORT is not set")
		return err
	}
	NewConfig.DBconn = fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable host=%s port=%s", NewConfig.DBuser, NewConfig.DBpassword, NewConfig.DBname, NewConfig.DBhost, NewConfig.DBport)
	return nil
}