package services

import (
	"errors"
	"mongo_data/core/data_generator"
	"mongo_data/core/data_operator"
	"mongo_data/core/parser_factory"
	"mongo_data/model"
)

func DataCreate(m *model.RequestVO) error {
	parser := parser_factory.GetParser(m.Type)
	if parser == nil {
		return errors.New("no parser error")
	}

	field, err := parser.Transfer(m.Content)
	if err != nil {
		return err
	}

	var dataList []interface{}

	generator := data_generator.GetGenerator(m.GenerateType)
	if generator == nil {
		return errors.New("no generator error")
	}
	for i := 0; i < m.Rows; i++ {
		data := make(map[string]interface{}, 0)
		data_generator.GenerateData(data, field, generator)
		dataList = append(dataList, data)
	}

	err = data_operator.InsertMany(m.DB, m.Col, dataList)
	return err

}
