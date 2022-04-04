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

func main() {
	conf, e := cfg.LoadConfiguration("C:\\configFile\\config.json")
	checkError(e)

	for i := 0; i < len(conf.ConfigsSql); i++ {
		db, err := dbConn.GetDbConnect(conf.ConfigsSql[i].Db)
		checkError(err)
		dbb, errr := db.ConnectingToDb(conf.ConfigsSql[i])
		checkError(errr)
		strArray := []string{"tables", "views", "procedures", "schemas"}

		for j := 0; j < len(strArray); j++ {
			switch strArray[j] {
			case "tables":
				arr1 := db.GetDDLTables(dbb)
				managingFiles.UnloadingTableDDl(arr1, conf.ConfigsSql[i].Dbname, strArray[j])
				remTable := managingFiles.RemoveTableFromLocal(conf.ConfigsSql[i].Dbname, strArray[j], arr1)
				git.CommitAndPush(remTable, conf.ConfigsSql[i].Dbname)
			case "views":
				arr2 := db.GetDDLViews(dbb)
				managingFiles.UnloadingTableDDl(arr2, conf.ConfigsSql[i].Dbname, strArray[j])
				remTable := managingFiles.RemoveTableFromLocal(conf.ConfigsSql[i].Dbname, strArray[j], arr2)
				git.CommitAndPush(remTable, conf.ConfigsSql[i].Dbname)
			case "procedures":
				arr3 := db.GetDDLProcedures(dbb)
				managingFiles.UnloadingTableDDl(arr3, conf.ConfigsSql[i].Dbname, strArray[j])
				remTable := managingFiles.RemoveTableFromLocal(conf.ConfigsSql[i].Dbname, strArray[j], arr3)
				git.CommitAndPush(remTable, conf.ConfigsSql[i].Dbname)
			case "schemas":
				arr4 := db.GetDDLSchemas(dbb)
				managingFiles.UnloadingTableDDl(arr4, conf.ConfigsSql[i].Dbname, strArray[j])
				remTable := managingFiles.RemoveTableFromLocal(conf.ConfigsSql[i].Dbname, strArray[j], arr4)
				git.CommitAndPush(remTable, conf.ConfigsSql[i].Dbname)
			default:
				fmt.Println("Something went wrong...")
			}
		}
	}

}

func checkError(err error) {
	if err != nil {
		fmt.Sprintf("Error %s", err)
		panic(err)
	}
}
