package dbConn

import (
	"fmt"
)

func GetDbConnect(dbtype string) (IDatabase, error) {
	switch dbtype {
	case "postgres":
		return &PostgresDb{}, nil
	case "mysql":
		return &MysqlDb{}, nil
	default:
		return nil, fmt.Errorf("wrong database type passed")
	}
}
