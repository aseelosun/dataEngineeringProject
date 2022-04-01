package dbConn

import (
	conf "dataEngineeringProject/config"
	"database/sql"
)

type IDatabase interface {
	ConnectingToDb(conf conf.SqlDbParams) (*sql.DB, error)
	GetDDLTables(db *sql.DB) []string
	//GetDDLViews(db *sql.DB) []string
	//GetDDLProcedures(db *sql.DB) []string
	//GetDDLSchemas(db *sql.DB) []string
	//UnloadingTableDDL(db *sql.DB)

}
