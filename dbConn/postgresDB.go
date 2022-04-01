package dbConn

import (
	conf "dataEngineeringProject/config"
	"database/sql"
	"fmt"
	"log"
	"time"
)

type PostgresDb struct {
}

func (p PostgresDb) ConnectingToDb(conf conf.SqlDbParams) (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", conf.Host, conf.Port, conf.User, conf.Password, conf.Dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}
	db.SetMaxIdleConns(conf.MaxIdleConns)
	db.SetMaxOpenConns(conf.MaxOpenConns)
	db.SetConnMaxLifetime(1800 * time.Second)
	fmt.Printf("Connected!\n")

	return db, err
}
func (p PostgresDb) GetDDLTables(db *sql.DB) []string {
	var (
		tableName   string
		tablesArray []string
	)

	rows, err := db.Query("SELECT table_name FROM information_schema.tables WHERE table_schema = 'public'")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&tableName)
		var tableDdl string
		rows, e4 := db.Query(`SELECT generate_create_table_statement($1)`, tableName)
		if e4 != nil {
			fmt.Println(e4)
		}
		for rows.Next() {
			e5 := rows.Scan(&tableDdl)
			if e5 != nil {
				fmt.Println(e5)
			}
		}
		tablesArray = append(tablesArray, tableName, tableDdl)
		if err != nil {
			panic(err)
		}
	}
	return tablesArray
}

//func (p PostgresDb) GetDDLViews(db *sql.DB) []string {
//	return nil
//}
//
//func (p PostgresDb) GetDDLProcedures(db *sql.DB) []string {
//	return nil
//}
//
//func (p PostgresDb) GetDDLSchemas(db *sql.DB) []string {
//	return nil
//}
