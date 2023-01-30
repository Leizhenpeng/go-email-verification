package models

type CommonResponse struct {
	Code int    `json:"Code" example:"200"`
	Msg  string `json:"Message" example:"Success"`
}

type PingResponse struct {
	CommonResponse
	Msg string `json:"Message" example:"pong"`
}
