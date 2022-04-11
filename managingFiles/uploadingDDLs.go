package managingFiles

import (
	"dataEngineeringProject/types"
	"io/ioutil"
	"os"
)

func UnloadingTableDDl(tableDdls []types.DataDDLs, cPath string, dbname string, tType string) error {
	e2 := os.MkdirAll(cPath+string(os.PathSeparator)+dbname+string(os.PathSeparator)+tType, 0755)
	tablesPath := cPath + string(os.PathSeparator) + dbname + string(os.PathSeparator) + tType + string(os.PathSeparator)
	if e2 != nil {
		return e2
	}
	for i := 0; i < len(tableDdls); i += 1 {
		file, err := os.Create(tablesPath + tableDdls[i].ObjectName + "_ddl.txt")
		if err != nil {
			return err
		}
		_, err2 := file.WriteString("\n" + tableDdls[i].ObjectDDL)
		if err2 != nil {
			return err2
		}

	}
	return nil
}

func tableInDb(tableName string, obj []types.DataDDLs) bool {
	for i := 0; i < len(obj); i++ {
		if (obj[i].ObjectName + "_ddl.txt") == tableName {
			return true
		}
	}
	return false
}

func RemoveTableFromLocal(cPath string, dbname string, folderName string, obj []types.DataDDLs) (string, error) {
	items, _ := ioutil.ReadDir(cPath + string(os.PathSeparator) + dbname + string(os.PathSeparator) + folderName)
	var removedFileName string
	for _, item := range items {
		if !tableInDb(item.Name(), obj) {
			removedFileName = item.Name()
			path := cPath + string(os.PathSeparator) + dbname + string(os.PathSeparator) + folderName + string(os.PathSeparator) + item.Name()
			errr := os.Remove(path)
			if errr != nil {
				return "", errr
			}
		}
	}
	return removedFileName, nil
}
