package server

type IndentityRequest struct {
	Name string `json:"name"`
}

type ListRequest struct {
	Name string `json:"name"`
}

type RelayRequest struct {
	Name    string      `query:"name"`
	Message interface{} `json:"message"`
}
