package main

import (
	conf "dataEngineeringProject/config"
	"database/sql"
	"log"
	"os"
)

type Ddls struct {
	ObjectName string
	ObjectDdl  string
}

func (p postgresFiles) unloadingTableDDL(conf conf.PostgresDb, db *sql.DB) (file *os.File) {
	var arrTables []string = p.getListOfAllTables(db)
	err := os.MkdirAll(conf.CatalogsPath+"\\"+conf.Dbname+"\\tables", 0755)
	tablesPath := conf.CatalogsPath + "\\" + conf.Dbname + "\\tables\\"
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

	}
	return file

}

func (m mysqlFiles) unloadingTableDDl(conf conf.MysqlDb, db *sql.DB) (file *os.File) {
	err := os.MkdirAll(conf.CatalogsPath+"\\"+conf.Dbname+"\\tables", 0755)
	tablesPath := conf.CatalogsPath + "\\" + conf.Dbname + "\\tables\\"
	if err != nil {
		log.Fatal(err)
	}
	var arrTables []string = m.getListOfAllTables(db)
	for i := 0; i < len(arrTables); i++ {
		file, err := os.Create(tablesPath + arrTables[i] + "_ddl.txt")
		if err != nil {
			log.Fatal(err)
		}

		var (
			tableName string
			tableDdl  string
		)

		rows, err := db.Query("SHOW CREATE TABLE " + arrTables[i])
		if err != nil {
			panic(err)
		}
		for rows.Next() {
			err := rows.Scan(&tableName, &tableDdl)
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
