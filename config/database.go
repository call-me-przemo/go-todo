package config

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

func Database() *sql.DB {

	user := GetEnv("MYSQL_USER")
	pass := GetEnv("MYSQL_PASSWORD")
	host := GetEnv("MYSQL_HOST")
	db := GetEnv("MYSQL_DATABASE")
	port := GetEnv("MYSQL_PORT")

	cfg := mysql.Config{
		User:   user,
		Passwd: pass,
		Addr:   host + ":" + port,
		DBName: db,
	}

	database, err := sql.Open("mysql", cfg.FormatDSN())

	if err != nil {
		Logger.Log(err)
		log.Fatal(err)
	} else {
		fmt.Println("Database Connection Successful")
	}

	_, err = database.Exec(`
		CREATE TABLE IF NOT EXISTS todos (
		    id INT AUTO_INCREMENT,
		    item TEXT NOT NULL,
		    completed BOOLEAN DEFAULT FALSE,
		    PRIMARY KEY (id)
		);
	`)

	if err != nil {
		fmt.Println(err)
	}

	return database
}
