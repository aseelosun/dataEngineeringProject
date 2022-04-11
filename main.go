package main

import (
	"dataEngineeringProject/chain"
	cfg "dataEngineeringProject/config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

func main() {
	conf, e := cfg.LoadConfiguration("C:\\configFile\\config.json")
	checkError(e)
	for i := 0; i < len(conf.ConfigsSql); i++ {
		strArray := []string{"tables", "views", "procedures", "schemas"}
		for j := 0; j < len(strArray); j++ {
			err := chain.ExecuteChain(strArray[j], i)
			if err != nil {
				_ = fmt.Sprintf("Error %s", err)
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
