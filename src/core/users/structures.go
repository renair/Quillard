package users

type User struct {
	Id          int64
	Email       string
	Password    string
	rawPassword string
	Nickname    string
}
