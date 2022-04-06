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

type Paths struct {
	CatalogsPath string `json:"catalogsPath"`
}

func LoadPaths(filename string) (Paths, error) {
	jsonFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	//fmt.Println(string(byteValue))
	var config Paths
	json.Unmarshal(byteValue, &config)
	fmt.Println(config)
	return config, err

}
