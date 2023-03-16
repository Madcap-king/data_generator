package parser_factory

import (
	"mongo_data/constants"
	"mongo_data/model"
)

type ParserType = int

type Parser interface {
	Transfer(string) (*model.Field, error)
}

var parserMap map[ParserType]Parser

func init() {
	parserMap = make(map[ParserType]Parser, 0)
	parserMap[constants.JSON] = &JsonParser{}
	parserMap[constants.XML] = &XmlParser{}
}

func GetParser(t ParserType) Parser {
	switch t {
	case constants.JSON:
		return parserMap[constants.JSON]
	}
	return nil
}
