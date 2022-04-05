package main

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"os"
)

var ErrorCanNotConnectDatabaseNoConfig = errors.New("can not connect to database: no database config provided")

type DBConfig struct {
	host     string
	port     string
	user     string
	password string
	dbname   string
}

func main() {
	err := loadGoDotEnv()
	if err != nil {
		panic(err)
	}

	dbConfig := getDatabaseConfig()
	if err != nil {
		panic(err)
	}

	db, err := connectToDB(dbConfig)
	if err != nil {
		panic(err)
	}
	defer db.Close()
}

func loadGoDotEnv() error {
	err := godotenv.Load(".env")
	if err != nil {
		return err
	}
	return nil
}

func getDatabaseConfig() *DBConfig {
	return &DBConfig{
		host:     os.Getenv("DB_HOST"),
		port:     os.Getenv("DB_PORT"),
		user:     os.Getenv("DB_USERNAME"),
		password: os.Getenv("DB_PASSWORD"),
		dbname:   os.Getenv("DB_NAME"),
	}
}

func connectToDB(config *DBConfig) (*sql.DB, error) {
	if config == nil {
		return nil, ErrorCanNotConnectDatabaseNoConfig
	}

	psqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.host, config.port, config.user, config.password, config.dbname,
	)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
