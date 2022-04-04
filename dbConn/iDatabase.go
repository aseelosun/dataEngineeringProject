package dbConn

import (
	conf "dataEngineeringProject/config"
	"dataEngineeringProject/types"
	"database/sql"
)

type IDatabase interface {
	ConnectingToDb(conf conf.SqlDbParams) (*sql.DB, error)
	GetDDLTables(db *sql.DB) ([]types.DataDDLs, error)
	GetDDLViews(db *sql.DB) ([]types.DataDDLs, error)
	GetDDLProcedures(db *sql.DB) ([]types.DataDDLs, error)
	GetDDLSchemas(db *sql.DB) ([]types.DataDDLs, error)
}
