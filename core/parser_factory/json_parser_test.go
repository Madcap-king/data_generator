package parser_factory

import (
	"log"
	"mongo_data/constants"
	"testing"
)

var content1 = "{\n    \"create_time\": \"date\", \n    \"nick_name\": \"string\", \n    \"info\": {\n        \"age\": \"int\"\n    }, \n    \"country\": [\n        \"string\", \n        \"string\"\n    ], \n    \"extra\": [\n        {\n            \"address\": \"string\"\n        }, \n        {\n            \"address\": \"string\"\n        }\n    ]\n}"
var content2 = "{\"country\": [\n        \"string\", \n        \"string\"\n    ]}"
var content3 = "{\"extra\": [\n        {\n            \"address\": \"string\"\n        }, \n        {\n            \"address\": \"string\"\n        }\n    ]}"
var content4 = "{\n    \"userName\": \"string\", \n    \"nick_name\": \"string\", \n    \"info\": {\n        \"age\": \"int64\"\n    }, \n    \"country\": [\n        \"string\", \n        \"string\"\n    ], \n    \"extra\": [\n        {\n            \"address\": [\n                \"string\", \n                \"string\", \n                \"string\"\n            ]\n        }, \n        {\n            \"address\": [\n                \"string\", \n                \"string\"\n            ]\n        }\n    ]\n}"
var content5 = "{ \"userName\": \"string\", \n    \"nick_name\": \"string\", \n    \"info\": {\n        \"age\": \"int\"\n    }}"
var content6 = "{\n   \"extra\":[\n      {\n         \"address\":[\"string\",\"string\",\"string\"]\n      }\n   ]\n}"

func TestJson(t *testing.T) {
	list := []string{content1, content2, content3, content4, content5, content6}
	parser := GetParser(constants.JSON)

	for _, item := range list {
		s, _ := parser.Transfer(item)
		log.Println(s)
	}

}
