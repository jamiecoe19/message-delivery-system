package message

type IdentityRequest struct {
	Name string `json:"name"`
}

type IndentityResponse struct {
	UserID uint64 `json:"userId"`
}

func (i IndentityResponse) Send() error {
	return nil
}
