package models

type ResponseModel struct {
	Data    map[string]interface{} `json:"data,omitempty"`
	Message string                 `json:"msg,omitempty"`
	Code    int                    `json:"code"`
}

type ResponseModel2 struct {
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"msg,omitempty"`
	Code    int         `json:"code"`
}
