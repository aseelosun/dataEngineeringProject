package main

import (
	conf "dataEngineeringProject/config"
	"database/sql"
)

type IDatabase interface {
	ConnectingToDb(conf conf.SqlDbParams) (*sql.DB, error)
	//GetListOfAllTables(db *sql.DB) ([]Ddls, error)
	//GetListOfAllViews(db *sql.DB) ([]string, error)
	//GetListOfAllProcedures(db *sql.DB) ([]string, error)
	//GetListOfAllSchemas(db *sql.DB) ([]string, error)
}
