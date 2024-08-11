package types

type Message struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

type ValidationError struct {
	StatusCode int      `json:"status_code"`
	Errors     []string `json:"errors"`
}
