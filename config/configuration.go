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
	Paths Path `json:"Paths"`
}
type Path struct {
	CatalogsPath string `json:"catalogsPath"`
}

func LoadPaths(filename string) (Paths, error) {
	jsonFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var config Paths
	err = json.Unmarshal(byteValue, &config)
	if err != nil {
		fmt.Println("Error during Unmarshal(): ", err)
	}
	return config, nil

}

type Github struct {
	Github GitConfig `json:"gitConf"`
}

type GitConfig struct {
	Username   string
	Password   string
	Repository string
	RemoteName string
}

func LoadGitConfigs(filename string) (Github, error) {
	jsonFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var config Github
	err = json.Unmarshal(byteValue, &config)
	if err != nil {
		fmt.Println("Error during Unmarshal(): ", err)
	}
	return config, nil

}
