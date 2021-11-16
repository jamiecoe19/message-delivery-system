package server

type IndentityRequest struct {
	Name string `json:"name"`
}

type ListRequest struct {
	Name string `json:"name"`
}

type RelayRequest struct {
	Sender     string      `json:"sender"`
	Recipients []string    `json:"recipients"`
	Message    interface{} `json:"message"`
}
