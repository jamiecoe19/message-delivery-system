package message

type RelayRequest struct {
	Name    string `json:"name"`
	Message string `Message`
}

type RelayResponse struct {
	Message interface{} `json: "message"`
}

func (response RelayResponse) Send() error {
	return nil
}
