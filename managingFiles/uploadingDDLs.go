package managingFiles

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func UnloadingTableDDl(tableDdls []string) {

	e2 := os.MkdirAll("C:\\Users\\Trainee\\dataEngineeringProject\\catalogs\\postgres\\tables", 0755)
	tablesPath := "C:\\Users\\Trainee\\dataEngineeringProject\\catalogs\\postgres\\tables\\"
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
