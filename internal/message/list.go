package message

type ListRequest struct {
	Name string `json:"name"`
}

type ListResponse struct {
	Users []uint64 `json:"users"`
}
