package positions

type Position struct {
	Id        int64   `json:"-"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
	Name      string  `json:"name,omitempty"`
}

//User request JSON objects
type PersonageId struct {
	Id int64 `json:"personage_id"`
}

//convertion for db query
func (pos Position) toKeys() map[string]interface{} {
	return map[string]interface{}{
		"longitude": pos.Longitude,
		"latitude":  pos.Latitude,
		"name":      pos.Name,
	}
}
