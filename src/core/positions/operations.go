package positions

//Get position information based on position's ID
func GetPositionById(id int64) *Position {
	checkConnection()
	keys := map[string]interface{}{
		"id": id,
	}
	result, dbErr := connection.SelectBy(TABLENAME, keys, "longitude", "latitude", "name")
	if dbErr != nil {
		return nil
	}
	pos := Position{
		Id: id,
	}
	result.Scan(&pos.Longitude, &pos.Latitude, &pos.Name)
	return &pos
}

//Select position information based on position's name
func GetPositionByName(name string) *Position {
	checkConnection()
	keys := map[string]interface{}{
		"name": name,
	}
	result, dbErr := connection.SelectBy(TABLENAME, keys, "id", "longitude", "latitude")
	if dbErr != nil {
		return nil
	}
	pos := Position{
		Name: name,
	}
	result.Scan(&pos.Id, &pos.Longitude, &pos.Latitude)
	return &pos
}

//Get position based on coords
func GetPositionByCoords(longitude float64, latitude float64) *Position {
	checkConnection()
	keys := map[string]interface{}{
		"longitude": longitude,
		"latitude":  latitude,
	}
	result, dbErr := connection.SelectBy(TABLENAME, keys, "id", "name")
	if dbErr != nil {
		return nil
	}
	pos := Position{
		Longitude: longitude,
		Latitude:  latitude,
	}
	result.Scan(&pos.Id, &pos.Name)
	return &pos
}

//Save position in db
func SavePosition(pos Position) error {
	checkConnection()
	return connection.Insert(TABLENAME, pos.toKeys())
}
