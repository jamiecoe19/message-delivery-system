package server

type User struct {
	UserID uint64 `json:"userId"`
	Name   string `json:"name"`
}
