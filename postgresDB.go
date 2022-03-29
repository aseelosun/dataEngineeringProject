package main

import (
	conf "dataEngineeringProject/config"
	"database/sql"
	"fmt"
	"log"
)

type PostgresDb struct {
	iDatabase
}

func (p PostgresDb) connectingTodDb(conf conf.PostgresDb) (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", conf.Host, conf.Port, conf.User, conf.Password, conf.Dbname)

	conn, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}
	fmt.Printf("Connected!\n")

	return conn, err
}

func (p PostgresDb) getListOfAllTables(db *sql.DB) []string {
	var (
		tableName string
		tableType string
		arrTables []string
	)

	rows, err := db.Query("show full tables where Table_Type = 'BASE TABLE'")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&tableName, &tableType)
		arrTables = append(arrTables, tableName)
		if err != nil {
			panic(err)
		}
	}
	return arrTables
}
