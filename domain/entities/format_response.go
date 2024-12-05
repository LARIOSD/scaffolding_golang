package entities

type ResponseHttp struct {
	Code    int    `json:"code"`
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type ResponseJson struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
