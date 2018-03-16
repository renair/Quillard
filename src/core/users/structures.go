package users

//DB Objects

//Type represents `user` table row
type User struct {
	Id          int64
	Email       string
	Password    string
	rawPassword string
	Nickname    string
}

//JSON Request Objects

type LoginRequest struct {
	Email    string `json:email`
	Password string `json:password`
}

type RegisterRequest struct {
	Nickname string `json:nickname`
	Email    string `json:email`
	Password string `json:password`
}
