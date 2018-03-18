package sessions

type Session struct {
	Id      int64
	UserId  int64
	Key     string
	Created int64
	Expires int64
}

//Clientside Session
type SessionResponse struct {
	Key     string `json:"key"`
	Expires int64  `json:"expires"`
}
