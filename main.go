package main

import (
	"dataEngineeringProject/chain"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

func main() {
	err := chain.ExecuteChain()
	checkError(err)

	//db, err := dbConn.GetDbConnect(conf.ConfigsSql[0].Db)
	//dbb, errr := db.ConnectingToDb(conf.ConfigsSql[0])
	//
	//checkError(e)
	//checkError(err)
	//checkError(errr)
	//
	//arr, e := db.GetDDLTables(dbb)
	//ee := managingFiles.UnloadingTableDDl(arr, cpath.Paths.CatalogsPath, conf.ConfigsSql[0].Dbname, "tables")
	//remf, e22 := managingFiles.RemoveTableFromLocal(cpath.Paths.CatalogsPath, conf.ConfigsSql[0].Dbname, "tables", arr)
	//fmt.Println(remf)
	//checkError(e22)
	//checkError(ee)

	//checkError(errr)

	//checkError(ee)

}

func checkError(err error) {
	if err != nil {
		fmt.Sprintf("Error %s", err)
		panic(err)
	}
}
