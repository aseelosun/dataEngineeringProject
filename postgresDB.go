package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type PostgresDbs struct {
	PostgresDbs []PostgresDb `json:"postgresDbs"`
}

type PostgresDb struct {
	Db           string
	Server       string
	Host         string
	Port         string
	User         string
	Password     string
	Dbname       string
	MaxIdleConns int
	MaxOpenConns int
	MaxLifetime  int
	CatalogsPath string
}

func loadConfiguration(filename string) (PostgresDbs, error) {
	jsonFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var configs PostgresDbs
	json.Unmarshal(byteValue, &configs)
	return configs, err
}

func connectingTodDb(p PostgresDb) (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", p.Host, p.Port, p.User, p.Password, p.Dbname)

	conn, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}
	fmt.Printf("Connected!\n")

	return conn, err

}

func getListOfAllTables(db *sql.DB) []string {
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

func unloadingTableDDL(db *sql.DB, p PostgresDb) (file *os.File) {
	var arrTables []string = getListOfAllTables(db)
	err := os.MkdirAll(p.CatalogsPath + "\\"
	p.Dbname + "\\tables", 0755)
	tablesPath := p.CatalogsPath + "\\" + p.Dbname + "\\tables\\"
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < len(arrTables); i++ {
		file, err := os.Create(tablesPath + arrTables[i] + "_ddl.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		var (
			tableDdl string
		)

		rows, err := db.Query(`SELECT generate_create_table_statement($1)`, arrTables[i])
		if err != nil {
			panic(err)
		}
		for rows.Next() {
			err := rows.Scan(&tableDdl)
			if err != nil {
				panic(err)
			}
			_, err2 := file.WriteString("\n" + tableDdl)
			if err2 != nil {
				log.Fatal(err2)
			}
		}
	}
	return file

}
