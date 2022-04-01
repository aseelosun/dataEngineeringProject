package dbConn

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

func (m MysqlDb) GetDDLTables(db *sql.DB) []string {
	var (
		tableName   string
		tableType   string
		tablesArray []string
	)

	rows, err := db.Query("show full tables where Table_Type = 'BASE TABLE'")
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&tableName, &tableType)
		if err != nil {
			fmt.Println(err)
		}
		var (
			tableNamee string
			tableDdl   string
		)

		rows, err := db.Query("SHOW CREATE TABLE " + tableName)
		if err != nil {
			panic(err)
		}
		if err != nil {
			panic(err)
		}
		for rows.Next() {
			err := rows.Scan(&tableNamee, &tableDdl)
			if err != nil {
				panic(err)
			}
			tablesArray = append(tablesArray, tableNamee, tableDdl)
		}
	}
	return tablesArray
}

//func (m MysqlDb) GetDDLViews(db *sql.DB) []string {
//	var (
//		tableName string
//		tableType string
//		arrTables []string
//	)
//
//	rows, err := db.Query("SHOW FULL TABLES IN mysql WHERE table_type LIKE 'VIEW'")
//	if err != nil {
//		panic(err)
//	}
//	defer rows.Close()
//	for rows.Next() {
//		err := rows.Scan(&tableName, &tableType)
//		if err != nil {
//			panic(err)
//		}
//		var (
//			viewName             string
//			viewDdl              string
//			character_set_client string
//			collation_connection string
//		)
//
//		rows, err := db.Query("SHOW CREATE VIEW " + tableName)
//		if err != nil {
//			panic(err)
//		}
//		for rows.Next() {
//			err := rows.Scan(&viewName, &viewDdl, &character_set_client, &collation_connection)
//			if err != nil {
//				panic(err)
//			}
//			arrTables = append(arrTables, tableName, viewDdl)
//		}
//	}
//	return arrTables
//}
//
//func (m MysqlDb) GetDDLProcedures(db *sql.DB) []string {
//	var (
//		dbb           string
//		name          string
//		ttype         string
//		definer       string
//		modified      string
//		created       string
//		sec_type      string
//		comment       string
//		ch_set        string
//		coll_conn     string
//		db_coll       string
//		arrProcedures []string
//	)
//
//	rows, err := db.Query("SHOW PROCEDURE STATUS")
//	if err != nil {
//		panic(err)
//	}
//	defer rows.Close()
//	for rows.Next() {
//		err := rows.Scan(&dbb, &name, &ttype, &definer, &modified, &created, &sec_type, &comment, &ch_set, &coll_conn, &db_coll)
//		if err != nil {
//			panic(err)
//		}
//		var (
//			procedure           string
//			sqlMode             string
//			createProcedure     string
//			characterSetClient  string
//			collationConnection string
//			databaseCollation   string
//		)
//
//		rows, err := db.Query("SHOW CREATE PROCEDURE " + name)
//		if err != nil {
//			panic(err)
//		}
//		defer rows.Close()
//		for rows.Next() {
//			err := rows.Scan(&procedure, &sqlMode, &createProcedure, &characterSetClient, &collationConnection, &databaseCollation)
//			if err != nil {
//				panic(err)
//			}
//			arrProcedures = append(arrProcedures, name, createProcedure)
//		}
//
//	}
//	return arrProcedures
//}
//
//func (m MysqlDb) GetDDlSchemas(db *sql.DB) []string {
//	var (
//		dbName     string
//		arrSchemas []string
//	)
//
//	rows, err := db.Query("show databases")
//	if err != nil {
//		panic(err)
//	}
//	for rows.Next() {
//		err := rows.Scan(&dbName)
//
//		if err != nil {
//			panic(err)
//		}
//		var (
//			database       string
//			createDatabase string
//		)
//
//		rows, err := db.Query("SHOW CREATE SCHEMA " + dbName)
//		if err != nil {
//			panic(err)
//		}
//
//		for rows.Next() {
//			err := rows.Scan(&database, &createDatabase)
//			if err != nil {
//				panic(err)
//			}
//			arrSchemas = append(arrSchemas, dbName, createDatabase)
//		}
//	}
//	return arrSchemas
//}
