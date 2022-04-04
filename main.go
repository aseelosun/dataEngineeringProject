package main

import (
	cfg "dataEngineeringProject/config"
	"dataEngineeringProject/dbConn"
	"dataEngineeringProject/git"
	"dataEngineeringProject/managingFiles"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

const (
	strArray = []string{"tables", "views", "procedures", "schemas"}
)

func main() {

	conf, e := cfg.LoadConfiguration("C:\\configFile\\config.json")
	checkError(e)

	for i := 0; i < len(conf.ConfigsSql); i++ {
		db, err := dbConn.GetDbConnect(conf.ConfigsSql[i].Db)

		if nil != err {
			checkError(err)
			continue
		}
		dbb, errr := db.ConnectingToDb(conf.ConfigsSql[i])
		checkError(errr)

		for strArray...)
		worker := new
		worker(strArray[i])
		if nil != err {
			checkError(err)
			continue
		}
		worker.GetDDLTables(dbb)
		worker.UnloadingTableDDl
		worker.CommitAndPush

		arr1, err := db.GetDDLTables(dbb)
		if nil != err {
			checkError(err)
			continue
		}
		err = managingFiles.UnloadingTableDDl(arr1, conf.ConfigsSql[i].Dbname, strArray[j])
		if nil != err {
			checkError(err)
			continue
		}
		remTable := managingFiles.RemoveTableFromLocal(conf.ConfigsSql[i].Dbname, strArray[j], arr1)
		if nil != err {
			checkError(err)
			continue
		}
		git.CommitAndPush(remTable, conf.ConfigsSql[i].Dbname)

		arr2 := db.GetDDLViews(dbb)
		managingFiles.UnloadingTableDDl(arr2, conf.ConfigsSql[i].Dbname, strArray[j])
		remTable := managingFiles.RemoveTableFromLocal(conf.ConfigsSql[i].Dbname, strArray[j], arr2)
		git.CommitAndPush(remTable, conf.ConfigsSql[i].Dbname)

		arr3 := db.GetDDLProcedures(dbb)
		managingFiles.UnloadingTableDDl(arr3, conf.ConfigsSql[i].Dbname, strArray[j])
		remTable := managingFiles.RemoveTableFromLocal(conf.ConfigsSql[i].Dbname, strArray[j], arr3)
		git.CommitAndPush(remTable, conf.ConfigsSql[i].Dbname)

		arr4 := db.GetDDLSchemas(dbb)
		managingFiles.UnloadingTableDDl(arr4, conf.ConfigsSql[i].Dbname, strArray[j])
		remTable := managingFiles.RemoveTableFromLocal(conf.ConfigsSql[i].Dbname, strArray[j], arr4)
		git.CommitAndPush(remTable, conf.ConfigsSql[i].Dbname)

	}
}

}

func checkError(err error) {
	if err != nil {
		fmt.Sprintf("Error %s", err)
		panic(err)
	}
}
