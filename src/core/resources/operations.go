package resources

import (
	"errors"
	"fmt"
)

var resourceTypes []Resource

func InitResourcesForPersonage(personageId int64) error {
	checkConnection()
	columns := []string{"resource_id", "personage_id", "amount"}
	data := make([][]interface{}, len(resourceTypes))
	for i := 0; i < len(data); i++ {
		data[i] = make([]interface{}, len(columns))
		data[i][0] = resourceTypes[i].Id
		data[i][1] = personageId
		data[i][2] = 0
	}
	return connection.MultipleInsert(STORAGETABLE, columns, data)
}

func AddResource(personageId int64, resourceId int32, amount int64) error {
	if !IsEnoughResource(personageId, resourceId, amount) {
		return errors.New("Not enough resources")
	}
	query := "UPDATE %s SET amount=amount-%d WHERE resource_id=%d AND personage_id=%d;"
	query = fmt.Sprintf(query, STORAGETABLE, amount, resourceId, personageId)
	_, err := connection.ManualQuery(query)
	return err
}

func GetPersonageResources(personageId int64) []ResourceResponse {
	query := `SELECT resource_id, name, amount FROM resource_storages INNER JOIN resource_types ON resource_storages.resource_id=resource_types.id AND personage_id=%d;`
	query = fmt.Sprintf(query, personageId)
	res, err := connection.ManualQuery(query)
	var result []ResourceResponse
	for err == nil && res.Next() {
		storedResource := ResourceResponse{}
		res.Scan(&storedResource.Id, &storedResource.Name, &storedResource.Amount)
		result = append(result, storedResource)
	}
	return result
}

func IsEnoughResource(personageId int64, resourceId int32, amount int64) bool {
	checkConnection()
	query := "SELECT * FROM %s WHERE resource_id=%d AND personage_id=%d AND amount >= %d;"
	query = fmt.Sprintf(query, STORAGETABLE, resourceId, personageId, amount)
	res, err := connection.ManualQuery(query)
	if err != nil {
		return false
	}
	defer res.Close()
	return res.Next()
}

//Must be called in init to init all resources list
func getResourceTypes() error {
	if len(resourceTypes) > 0 {
		return nil
	}
	checkConnection()
	res, err := connection.Select(TYPETABLE, "id", "name")
	if err != nil {
		return err
	}
	defer res.Close()
	for res.Next() {
		resType := Resource{}
		res.Scan(&resType.Id, &resType.Name)
		resourceTypes = append(resourceTypes, resType)
	}
	return nil
}
