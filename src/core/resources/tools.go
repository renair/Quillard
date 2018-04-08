package resources

func checkConnection() {
	if connection == nil {
		panic("Connection is not initialized in resouces package")
	}
}
