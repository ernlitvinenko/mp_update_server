package responses

type ErrorResponse struct {
	Status int    `json:"status"`
	Error  string `json:"error"`
}
