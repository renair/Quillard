package positions

type Position struct {
	Id        int64   `json:"-"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitide"`
	Name      string  `json:"name,omitempty"`
}

func (pos Position) toKeys() map[string]interface{} {
	return map[string]interface{}{
		"id":        pos.Id,
		"longitude": pos.Longitude,
		"latitude":  pos.Latitude,
		"name":      pos.Name,
	}
}
