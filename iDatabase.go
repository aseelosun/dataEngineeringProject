package main

type iDatabase interface {
	connectingToDb()
	getListOfAllTables()
	getListOfAllViews()
	getListOfAllProcedures()
	getListOfAllSchemas()
}
