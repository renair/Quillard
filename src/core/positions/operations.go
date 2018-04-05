package positions

import (
	"fmt"
)

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
	defer result.Close()
	pos := Position{
		Id: id,
	}
	result.Next()
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
	defer result.Close()
	pos := Position{
		Name: name,
	}
	result.Next()
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
	defer result.Close()
	pos := Position{
		Longitude: longitude,
		Latitude:  latitude,
	}
	result.Next()
	result.Scan(&pos.Id, &pos.Name)
	return &pos
}

//return position where account home set
func GetAccountHomePosition(accountId int64) *Position {
	query := `SELECT id, longitude, latitude, name from %s WHERE id IN (
		SELECT home_position FROM accounts WHERE id=%d
	);`
	query = fmt.Sprintf(query, TABLENAME, accountId)
	res, err := connection.ManualQuery(query)
	if err == nil && res.Next() {
		defer res.Close()
		pos := Position{}
		res.Scan(&pos.Id, &pos.Longitude, &pos.Latitude, &pos.Name)
		return &pos
	} else {
		return nil
	}
}

func CanBuild(pos Position) bool {
	query := "SELECT * FROM positions WHERE (latitude BETWEEN %f AND %f) AND (longitude BETWEEN %f AND %f)"
	query = fmt.Sprintf(query, pos.Latitude-BUILDDISTANCE, pos.Latitude+BUILDDISTANCE, pos.Longitude-BUILDDISTANCE, pos.Longitude+BUILDDISTANCE)
	res, err := connection.ManualQuery(query)
	return err == nil && !res.Next()
}

//Save position in db
func SavePosition(pos Position) *Position {
	checkConnection()
	connection.BeginTransaction()
	if connection.Insert(TABLENAME, pos.toKeys()) != nil {
		connection.RollbackTransaction()
		return nil
	} else {
		connection.CommitTransaction()
		return GetPositionByCoords(pos.Longitude, pos.Latitude)
	}
}
