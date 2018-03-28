package basehandlers

//Error structure
type StatusedResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
