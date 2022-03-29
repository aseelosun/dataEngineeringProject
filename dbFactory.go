package main

func doProcess(dbtype string) (iDatabase, error) {
	switch dbtype {
	case "postgres":
		return
	case "mysql":
		return
	default:
		return nil
	}

}
