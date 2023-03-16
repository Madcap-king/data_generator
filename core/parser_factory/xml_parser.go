package parser_factory

import "mongo_data/model"

type XmlParser struct {
}

func (p *XmlParser) Transfer(string) (*model.Field, error) {
	// todo
	return &model.Field{}, nil
}
