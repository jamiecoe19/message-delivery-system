package helpers

type RelayRequest struct {
	Sender     string      `json:"sender"`
	Recipients []string    `json:"recipients"`
	Message    interface{} `json:"message"`
}
