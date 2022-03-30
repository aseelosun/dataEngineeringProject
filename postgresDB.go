package main

import (
	conf "dataEngineeringProject/config"
	"database/sql"
	"fmt"
	"log"
)

type PostgresDb struct {
}

func (p PostgresDb) ConnectingToDb(conf conf.SqlDbParams) (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", conf.Host, conf.Port, conf.User, conf.Password, conf.Dbname)

	conn, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}
	fmt.Printf("Connected!\n")

	return conn, err
}

//func getDdlTable(tableName) (ddl string) {
//	rows, err := db.Query(`SELECT generate_create_table_statement($1)`, arrTables[i])
//	if err != nil {
//		panic(err)
//	}
//	for rows.Next() {
//		err := rows.Scan(&tableDdl)
//		if err != nil {
//			panic(err)
//		}
//		_, err2 := file.WriteString("\n" + tableDdl)
//		if err2 != nil {
//			log.Fatal(err2)
//		}
//	}
//}
//func (p PostgresDb) GetListOfAllTables(db *sql.DB) ([]Ddls, error) {
//	var (
//		tableName string
//		tableType string
//		arrTables []Ddls
//	)
//
//	rows, err := db.Query("show full tables where Table_Type = 'BASE TABLE'")
//	if err != nil {
//		panic(err)
//	}
//	defer rows.Close()
//	for rows.Next() {
//		err := rows.Scan(&tableName, &tableType)
//		if err != nil {
//			return nil, err
//		}
//		ddl, err := getDdlTable(tableName)
//		arrTables = append(arrTables, Ddls{tableName, ddl})
//	}
//	return arrTables, nil
//}
//func (p PostgresDb) GetListOfAllViews(db *sql.DB) ([]string, error) {
//	return nil, fmt.Errorf("www")
//}
//func (p PostgresDb) GetListOfAllProcedures(db *sql.DB) ([]string, error) {
//	return nil, fmt.Errorf("www")
//}
//func (p PostgresDb) GetListOfAllSchemas(db *sql.DB) ([]string, error) {
//	return nil, fmt.Errorf("www")
//}
