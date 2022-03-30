package main

import (
	cfg "dataEngineeringProject/config"
	"fmt"
	_ "github.com/lib/pq"
)

func main() {

	conf, e := cfg.LoadConfiguration("C:\\configFile\\config.json")

	for i := 0; i < len(conf.ConfigsSql); i++ {
		fmt.Println(conf.ConfigsSql[i].Db)
		db, err := GetDbConnect("postgres")
		fmt.Println(db)
		checkError(err)
		db.ConnectingToDb(conf.ConfigsSql[1])

	}

	checkError(e)

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

func checkError(err error) {
	if err != nil {
		fmt.Sprintf("Error %s", err)
		panic(err)
	}

}
