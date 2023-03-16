package test

import (
	"mongo_data/constants"
	"mongo_data/core/data_generator"
	"mongo_data/core/data_operator"
	"mongo_data/core/parser_factory"
	"testing"
)

var content1 = "{\n    \"create_time\": \"date\", \n    \"nick_name\": \"string\", \n    \"info\": {\n        \"age\": \"int32\"\n    }, \n    \"country\": [\n        \"string\", \n        \"string\"\n    ], \n    \"extra\": [\n        {\n            \"address\": \"string\"\n        }, \n        {\n            \"address\": \"string\"\n        }\n    ]\n}"
var content2 = "{\"country\": [\n        \"string\", \n        \"string\"\n    ]}"
var content3 = "{\"extra\": [\n        {\n            \"address\": \"string\"\n        }, \n        {\n            \"address\": \"string\"\n        }\n    ]}"
var content4 = "{\n    \"userName\": \"string\", \n    \"nick_name\": \"string\", \n    \"info\": {\n        \"age\": \"int64\"\n    }, \n    \"country\": [\n        \"string\", \n        \"string\"\n    ], \n    \"extra\": [\n        {\n            \"address\": [\n                \"string\", \n                \"string\", \n                \"string\"\n            ]\n        }, \n        {\n            \"address\": [\n                \"string\", \n                \"string\"\n            ]\n        }\n    ]\n}"
var content5 = "{ \"userName\": \"string\", \n    \"nick_name\": \"string\", \n    \"info\": {\n        \"age\": \"int32\"\n    }}"
var content6 = "{\n   \"extra\":[\n      {\n         \"address\":[\"string\",\"string\",\"string\"]\n      }\n   ]\n}"

func TestJson(t *testing.T) {
	data_operator.MongoUri = "mongodb://127.0.0.1:27017"
	list := []string{content1, content2, content3, content4, content5, content6}
	parser := parser_factory.GetParser(constants.JSON)
	var dataList []interface{}
	generator := data_generator.GetGenerator(constants.Random)
	for _, item := range list {
		s, _ := parser.Transfer(item)
		data := make(map[string]interface{}, 0)
		data_generator.GenerateData(data, s, generator)
		dataList = append(dataList, data)
	}
	data_operator.InsertMany("test10", "test10", dataList)
}
