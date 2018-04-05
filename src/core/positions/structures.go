package positions

type Position struct {
	Id        int64   `json:"-"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
	Name      string  `json:"name,omitempty"`
}

//convertion for db query
func (pos Position) toKeys() map[string]interface{} {
	return map[string]interface{}{
		"longitude": pos.Longitude,
		"latitude":  pos.Latitude,
		"name":      pos.Name,
	}
}
