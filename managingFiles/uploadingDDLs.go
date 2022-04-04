package managingFiles

import (
	"dataEngineeringProject/types"
	"io/ioutil"
	"log"
	"os"
)

func UnloadingTableDDl(tableDdls []types.DataDDLs, dbname string, tType string) {

	e2 := os.MkdirAll("C:\\Users\\Trainee\\dataEngineeringProject\\catalogs\\"+dbname+"\\"+tType, 0755)
	tablesPath := "C:\\Users\\Trainee\\dataEngineeringProject\\catalogs\\" + dbname + "\\" + tType + "\\"
	if e2 != nil {
		panic(e2)
	}
	for i := 0; i < len(tableDdls); i += 1 {
		file, err := os.Create(tablesPath + tableDdls[i].ObjectName + "_ddl.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		_, err2 := file.WriteString("\n" + tableDdls[i].ObjectDDL)
		if err2 != nil {
			log.Fatal(err2)
		}

	}
}

func tableInDb(tableName string, tablesList []string) bool {
	for i := 0; i < len(tablesList); i += 2 {
		if (tablesList[i] + "_ddl.txt") == tableName {
			return true
		}
	}
	return false
}

func RemoveTableFromLocal(dbname string, folderName string, arr []string) string {
	items, _ := ioutil.ReadDir("C:\\Users\\Trainee\\dataEngineeringProject\\catalogs\\" + dbname + "\\" + folderName)
	var removedFileName string
	for _, item := range items {
		if !tableInDb(item.Name(), arr) {
			removedFileName = item.Name()
			path := "C:\\Users\\Trainee\\dataEngineeringProject\\catalogs\\" + dbname + "\\" + folderName + "\\" + item.Name()
			errr := os.Remove(path)
			if errr != nil {
				panic(errr)
			}
		}
	}
	return removedFileName
}
