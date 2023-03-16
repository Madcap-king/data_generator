package model

type RequestVO struct {
	DB           string `json:"db"`
	Col          string `json:"col"`
	Content      string `json:"content"`
	Type         int    `json:"type"`
	Rows         int    `json:"rows"`
	GenerateType int    `json:"generate_type"`
}

type ResponseVO struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
