package personages

import (
	"core/sessions"
	"fmt"
)

func registerPersonage(ses *sessions.Session, req PersonageRequest) error {
	checkConnection()
	keys := map[string]interface{}{
		"account_id":  ses.AccountId,
		"name":        req.Name,
		"position_id": 1, //TODO replace with home position id
	}
	return connection.Insert(TABLENAME, keys)
}

func getAccountPersonages(ses *sessions.Session) []PersonageResponse {
	checkConnection()
	res := make([]PersonageResponse, 0)
	query := `select personages.id, personages.name, positions.latitude, positions.longitude, positions.name 
			  from personages inner join positions 
			  on personages.position_id=positions.id
			  where personages.account_id=%d;`
	query = fmt.Sprintf(query, ses.AccountId)
	data, selectErr := connection.ManualQuery(query)
	if selectErr == nil {
		defer data.Close()
		for data.Next() {
			personage := PersonageResponse{}
			data.Scan(&personage.Id, &personage.Name, &personage.Position.Latitude, &personage.Position.Longitude, &personage.Position.Name)
			res = append(res, personage)
		}
	}
	return res
}
