package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type ConfigsMysql struct {
	ConfigsMysql []MysqlDb `json:"mysqlDbs"`
}
type ConfigsPostgres struct {
	ConfigsPostgres []PostgresDb `json:"postgresDbs"`
}

type PostgresDb struct {
	Db           string
	Server       string
	Host         string
	Port         string
	User         string
	Password     string
	Dbname       string
	CatalogsPath string
}

type MysqlDb struct {
	Db            string
	Server        string
	Host          string
	Port          string
	User          string
	Password      string
	Dbname        string
	MaxIdleConns  int
	MaxOpenConns  int
	MaxLifetime   int
	CatalogsPath  string
	TablesPath    string
	ViewsPath     string
	ProcedurePath string
	SchemaPath    string
}

func (p PostgresDb) loadConfiguration(filename string) (ConfigsPostgres, error) {
	jsonFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var configs ConfigsPostgres
	json.Unmarshal(byteValue, &configs)
	return configs, err
}

func (m MysqlDb) loadConfiguration(filename string) (ConfigsMysql, error) {
	jsonFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var configs ConfigsMysql

	json.Unmarshal(byteValue, &configs)
	return configs, err

}
