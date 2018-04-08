package resources

// DB types

type Resource struct {
	Id   int32
	Name string
}

type ResourceStorage struct {
	Id          int64
	ResourceId  int32
	PersonageId int64
	Amount      int64
}

// JSON to response

type ResourceResponse struct {
	Id     int64  `json:"id"`
	Name   string `json:"name"`
	Amount int64  `json:"amount"`
}

// requests JSON

type ResourceRequest struct {
	PersonageId int64 `json:"personage_id"`
}
