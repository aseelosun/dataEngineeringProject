package chain

import (
	cfg "dataEngineeringProject/config"
	"dataEngineeringProject/dbConn"
	"dataEngineeringProject/git"
	"dataEngineeringProject/managingFiles"
	"dataEngineeringProject/types"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

type ParameterValues struct {
	d           dbConn.IDatabase
	dbConn      *sql.DB
	dbname      string
	tType       string
	tableDdls   []types.DataDDLs
	removedFile string
}

type Chain interface {
	Action(p *ParameterValues) ([]types.DataDDLs, string, error)
}

type ddlGet struct {
}

func (d *ddlGet) Action(p *ParameterValues) ([]types.DataDDLs, string, error) {
	switch p.tType {
	case "tables":
		arr, e := p.d.GetDDLTables(p.dbConn)
		if e != nil {
			fmt.Sprintf("Error %s", e)
			return arr, "", e
		}
		fmt.Println("get table ddl done")
		return arr, "", nil

	case "views":
		arr, e := p.d.GetDDLViews(p.dbConn)
		if e != nil {
			fmt.Sprintf("Error %s", e)
			return arr, "", e
		}
		fmt.Println("get views ddl done")
		return arr, "", nil

	case "procedures":
		arr, e := p.d.GetDDLProcedures(p.dbConn)
		if e != nil {
			fmt.Sprintf("Error %s", e)
			return arr, "", e
		}
		fmt.Println("get procedures ddl done")
		return arr, "", nil

	case "schemas":
		arr, e := p.d.GetDDLSchemas(p.dbConn)
		if e != nil {
			fmt.Sprintf("Error %s", e)
			return arr, "", e
		}
		fmt.Println("get schemas ddl done")
		return arr, "", nil
	default:
		fmt.Println("Cannot get ddl")
	}
	var arr []types.DataDDLs
	return arr, "", nil

}

type ddlUpload struct {
	caller Chain
}

func (d *ddlUpload) Action(p *ParameterValues) ([]types.DataDDLs, string, error) {
	var arr []types.DataDDLs
	err := managingFiles.UnloadingTableDDl(p.tableDdls, p.dbname, p.tType)
	if err != nil {
		fmt.Sprintf("Error %s", err)
		return arr, "", err
	}
	fmt.Println("unloading ddl done")
	return arr, "", nil
}

type ddlRemove struct {
	caller Chain
}

func (d *ddlRemove) Action(p *ParameterValues) ([]types.DataDDLs, string, error) {
	removedFile, err := managingFiles.RemoveTableFromLocal(p.dbname, p.tType, p.tableDdls)
	var arr []types.DataDDLs
	if err != nil {
		fmt.Sprintf("Error %s", err)
		return arr, "", err
	}
	fmt.Println("removing ddl done")
	return arr, removedFile, nil
}

type ddlCommit struct {
	caller Chain
}

func (d *ddlCommit) Action(p *ParameterValues) ([]types.DataDDLs, string, error) {
	err := git.CommitAndPush(p.removedFile, p.dbname)
	var arr []types.DataDDLs
	if err != nil {
		fmt.Sprintf("Error %s", err)
		return arr, "", err
	}
	fmt.Println("removing ddl done")
	return arr, "", nil
}

func ExecuteChain() error {
	conf, e := cfg.LoadConfiguration("C:\\configFile\\config.json")
	checkError(e)

	for i := 0; i < len(conf.ConfigsSql); i++ {
		d, err := dbConn.GetDbConnect(conf.ConfigsSql[i].Db)
		checkError(err)
		db, err1 := d.ConnectingToDb(conf.ConfigsSql[i])
		checkError(err1)
		strArray := []string{"tables", "views", "procedures", "schemas"}
		for j := 0; j < len(strArray); j++ {
			prm := ParameterValues{
				d:      d,
				dbConn: db,
				dbname: conf.ConfigsSql[i].Dbname,
				tType:  strArray[j],
			}

			chain0 := &ddlGet{}

			chain1 := &ddlUpload{
				caller: chain0,
			}
			chain2 := &ddlRemove{
				caller: chain1,
			}

			finalChain := &ddlCommit{caller: chain2}

			_, _, err := finalChain.Action(&prm)

			if err != nil {
				fmt.Sprintf("Error %s", err)
				return err
			}

			fmt.Println("Success!")

			return nil
		}
	}
	return nil
}

func checkError(err error) {
	if err != nil {
		fmt.Sprintf("Error %s", err)
		panic(err)
	}
}
