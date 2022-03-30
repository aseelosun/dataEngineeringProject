package main

import (
	conf "dataEngineeringProject/config"
	"database/sql"
	"fmt"
	"time"
)

type MysqlDb struct {
}

func (m MysqlDb) ConnectingToDb(conf conf.SqlDbParams) (*sql.DB, error) {
	mysqlConn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		conf.User, conf.Password, conf.Host, conf.Port, conf.Dbname)

	db, err := sql.Open("mysql", mysqlConn)
	db.SetMaxIdleConns(conf.MaxIdleConns)
	db.SetMaxOpenConns(conf.MaxOpenConns)
	db.SetConnMaxLifetime(1800 * time.Second)
	fmt.Printf("Connected!\n")

	return db, err
}

//func (m MysqlDb) GetListOfAllTables(db *sql.DB) ([]string, error) {
//	var (
//		tableName string
//		tableType string
//		arrTables []string
//	)
//
//	rows, err := db.Query("show full tables where Table_Type = 'BASE TABLE'")
//	if err != nil {
//		panic(err)
//	}
//	defer rows.Close()
//	for rows.Next() {
//		err := rows.Scan(&tableName, &tableType)
//		arrTables = append(arrTables, tableName)
//		if err != nil {
//			return nil, err
//		}
//	}
//	return arrTables, nil
//}
//func (m MysqlDb) GetListOfAllViews(db *sql.DB) ([]string, error) {
//	return nil, fmt.Errorf("www")
//}
//func (m MysqlDb) GetListOfAllProcedures(db *sql.DB) ([]string, error) {
//	return nil, fmt.Errorf("www")
//}
//func (m MysqlDb) GetListOfAllSchemas(db *sql.DB) ([]string, error) {
//	return nil, fmt.Errorf("www")
//}
