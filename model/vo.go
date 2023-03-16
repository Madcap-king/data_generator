package model

type RequestVO struct {
	DB           string `json:"db"`            // 数据库名称
	Col          string `json:"col"`           // 集合名称
	Content      string `json:"content"`       // 文档字段名+字段属性
	Type         int    `json:"type"`          // 数据格式：目前支持content为json格式
	Rows         int    `json:"rows"`          // 需要模拟生成的文档数量
	GenerateType int    `json:"generate_type"` // 模拟数据的生成规则：random、词库， 目前只支持random
}

type ResponseVO struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
