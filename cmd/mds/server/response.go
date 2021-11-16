package server

type User struct {
	UserID uint64 `json:"userId"`
	Name   string `json:"name"`
}

type Success struct {
	Success bool `json:"success"`
}

func NewSuccessfulResponse() Success {
	return Success{
		Success: true,
	}
}
