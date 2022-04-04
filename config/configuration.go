package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type ConfigsSql struct {
	ConfigsSql []SqlDbParams `json:"databases"`
}

type Paths struct {
	CatalogsPath string
}

type SqlDbParams struct {
	Db           string
	Server       string
	Host         string
	Port         string
	User         string
	Password     string
	Dbname       string
	MaxIdleConns int
	MaxOpenConns int
	MaxLifetime  int
}

func LoadConfiguration(filename string) (ConfigsSql, error) {
	jsonFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var configs ConfigsSql
	json.Unmarshal(byteValue, &configs)
	return configs, err
}
