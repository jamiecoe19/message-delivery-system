package message

type Users []User

type User struct {
	UserID uint64 `json:"userId"`
	Name   string `json:"name"`
}
