package main

type iDatabase interface {
	loadConfiguration(db string, filename string)
	connectingToDb()
	getListOfAllTables()
	getListOfAllViews()
	getListOfAllProcedures()
	getListOfAllSchemas()
	unloadingTableDDL()
	unloadingViewDDL()
	unloadingProcedureDDL()
	unloadingSchemaDDL()
	removeFromLocal()
	gitCommitAndPush()
}
