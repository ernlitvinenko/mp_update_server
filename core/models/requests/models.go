package requests

type CreateApplicationRequest struct {
	Id          string  `json:"id"`
	AppName     string  `json:"appName"`
	Link        string  `json:"link"`
	Version     string  `json:"version"`
	Description *string `json:"description,omitempty"`
}

type AddVersionRequest struct {
	Id          string `json:"id"`
	AppId       string `json:"appId"`
	Description string `json:"description"`
	Link        string `json:"link"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
