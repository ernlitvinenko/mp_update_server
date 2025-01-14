package requests

type CreateApplicationRequest struct {
	Id      string `json:"id"`
	AppName string `json:"appName"`
	Link    string `json:"link"`
}

type AddVersionRequest struct {
	Id    string `json:"id"`
	AppId string `json:"app_id"`
	Link  string `json:"link"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
