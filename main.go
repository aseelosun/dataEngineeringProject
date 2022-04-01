package main

import (
	cfg "dataEngineeringProject/config"
	"dataEngineeringProject/dbConn"
	"dataEngineeringProject/git"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	conf, e := cfg.LoadConfiguration("C:\\configFile\\config.json")

	db, err := dbConn.GetDbConnect(conf.ConfigsSql[0].Db)
	dbb, errr := db.ConnectingToDb(conf.ConfigsSql[0])
	checkError(errr)

	//for i := 0; i < len(conf.ConfigsSql); i++ {
	//	fmt.Println(conf.ConfigsSql[i].Db)
	//	db, err := dbConn.GetDbConnect(conf.ConfigsSql[i].Db)
	//	fmt.Println(db)
	//	checkError(err)
	//	db.ConnectingToDb(conf.ConfigsSql[1])
	//
	//}
	checkError(e)
	checkError(err)
	newV := GetDDLTables(dbb)
	//fmt.Println(newV[0])
	UnloadingTableDDl(newV)

	remfile := RemoveTableFromLocal("tables", newV)
	fmt.Println(remfile)
	git.CommitAndPush(remfile)

	//config
	//for (all
	//conf
	//blocks){
	//	db, err := GetDbConnect("postgres")
	//	if nil != err {
	//		fmt.Errorf("asfjbasfjb")
	//		return
	//	}
	//
	//	db.ConnectingToDb()
	//	allDddlTables := db.GetListOfAllTables()
	//
	//	storage = fabric.GetStrorge(cfg)
	//	storage.saveTables(allDddlTables)
	//
	//}
	//commit
	//puch
}

//
func checkError(err error) {
	if err != nil {
		fmt.Sprintf("Error %s", err)
		panic(err)
	}
}

func UnloadingTableDDl(tableDdls []string) {

	e2 := os.MkdirAll("C:\\Users\\Trainee\\dataEngineeringProject\\catalogs\\mysql\\tables", 0755)
	tablesPath := "C:\\Users\\Trainee\\dataEngineeringProject\\catalogs\\mysql\\tables\\"
	if e2 != nil {
		fmt.Println(e2)
	}
	for i := 0; i < len(tableDdls); i += 2 {
		file, err := os.Create(tablesPath + tableDdls[i] + "_ddl.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		_, err2 := file.WriteString("\n" + tableDdls[i+1])
		if err2 != nil {
			log.Fatal(err2)
		}

	}
}
func GetDDLTables(db *sql.DB) []string {
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

func tableInDb(tableName string, tablesList []string) bool {
	for i := 0; i < len(tablesList); i += 2 {
		if (tablesList[i] + "_ddl.txt") == tableName {
			return true
		}
	}
	return false
}

func RemoveTableFromLocal(folderName string, arr []string) string {
	items, _ := ioutil.ReadDir("C:\\Users\\Trainee\\dataEngineeringProject\\catalogs\\mysql\\" + folderName)
	var removedFileName string
	for _, item := range items {
		if !tableInDb(item.Name(), arr) {
			removedFileName = item.Name()
			path := "C:\\Users\\Trainee\\dataEngineeringProject\\catalogs\\mysql\\" + folderName + "\\" + item.Name()
			errr := os.Remove(path)
			if errr != nil {
				fmt.Println(errr)
			}
		}
	}
	return removedFileName
}
