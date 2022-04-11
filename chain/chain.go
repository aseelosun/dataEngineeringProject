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

const ConfigPath = "C:\\configFile\\config.json"

type ParameterValues struct {
	d             dbConn.IDatabase
	dbConn        *sql.DB
	cPath         string
	dbname        string
	tType         string
	gitUsername   string
	gitPassword   string
	gitRepo       string
	gitRemoteName string
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
		fmt.Println("getting tables ddl is done")
		return arr, "", nil

	case "views":
		arr, e := p.d.GetDDLViews(p.dbConn)
		if e != nil {
			fmt.Sprintf("Error %s", e)
			return arr, "", e
		}
		fmt.Println("getting views ddl is done")
		return arr, "", nil

	case "procedures":
		arr, e := p.d.GetDDLProcedures(p.dbConn)
		if e != nil {
			fmt.Sprintf("Error %s", e)
			return arr, "", e
		}
		fmt.Println("getting procedures is done")
		return arr, "", nil

	case "schemas":
		arr, e := p.d.GetDDLSchemas(p.dbConn)
		if e != nil {
			fmt.Sprintf("Error %s", e)
			return arr, "", e
		}
		fmt.Println("getting schemas ddl is done")
		return arr, "", nil
	default:
		fmt.Println("Cannot get ddl")
	}
	var arr []types.DataDDLs
	return arr, "", nil

}

type ddlUpload struct {
	tableDdls []types.DataDDLs
	caller    Chain
}

func (d *ddlUpload) Action(p *ParameterValues) ([]types.DataDDLs, string, error) {
	var arr []types.DataDDLs
	err := managingFiles.UnloadingTableDDl(d.tableDdls, p.cPath, p.dbname, p.tType)
	if err != nil {
		fmt.Sprintf("Error %s", err)
		return arr, "", err
	}
	fmt.Println("unloading ddl is done")
	return arr, "", nil
}

type ddlRemove struct {
	tableDdls []types.DataDDLs
	caller    Chain
}

func (d *ddlRemove) Action(p *ParameterValues) ([]types.DataDDLs, string, error) {
	removedFile, err := managingFiles.RemoveTableFromLocal(p.cPath, p.dbname, p.tType, d.tableDdls)
	var arr []types.DataDDLs
	if err != nil {
		fmt.Sprintf("Error %s", err)
		return arr, "", err
	}
	fmt.Println("removing ddl is done")
	return arr, removedFile, nil
}

type ddlCommit struct {
	removedFile string
	caller      Chain
}

func (d *ddlCommit) Action(p *ParameterValues) ([]types.DataDDLs, string, error) {
	err := git.CommitAndPush(p.gitRemoteName, p.gitUsername, p.gitPassword, p.gitRepo, p.cPath, d.removedFile, p.dbname)
	var arr []types.DataDDLs
	if err != nil {
		fmt.Sprintf("Error %s", err)
		return arr, "", err
	}
	fmt.Println("committing ddl is done")
	return arr, "", nil
}

func ExecuteChain(tType string, i int) error {
	conf, e := cfg.LoadConfiguration(ConfigPath)
	checkError(e)
	paths, ee := cfg.LoadPaths(ConfigPath)
	checkError(ee)
	gitConf, eee := cfg.LoadGitConfigs(ConfigPath)
	checkError(eee)

	d, err := dbConn.GetDbConnect(conf.ConfigsSql[i].Db)
	checkError(err)
	db, err1 := d.ConnectingToDb(conf.ConfigsSql[i])
	checkError(err1)
	prm := ParameterValues{
		d:           d,
		dbConn:      db,
		cPath:       paths.Paths.CatalogsPath,
		dbname:      conf.ConfigsSql[i].Dbname,
		tType:       tType,
		gitUsername: gitConf.Github.Username,
		gitPassword: gitConf.Github.Password,
		gitRepo:     gitConf.Github.Repository,
	}

	chain0 := &ddlGet{}

	arr, _, e1 := chain0.Action(&prm)
	checkError(e1)

	chain1 := &ddlUpload{
		tableDdls: arr,
		caller:    chain0,
	}
	_, _, e2 := chain1.Action(&prm)
	checkError(e2)

	chain2 := &ddlRemove{
		tableDdls: arr,
		caller:    chain1,
	}
	_, remFile, e3 := chain2.Action(&prm)
	checkError(e3)

	finalChain := &ddlCommit{removedFile: remFile, caller: chain2}

	_, _, e4 := finalChain.Action(&prm)
	checkError(e4)

	fmt.Println("Success!")

	return nil

}

func checkError(err error) {
	if err != nil {
		fmt.Sprintf("Error %s", err)
	}
}
