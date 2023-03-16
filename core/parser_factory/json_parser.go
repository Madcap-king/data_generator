package parser_factory

import (
	"encoding/json"
	"mongo_data/constants"
	"mongo_data/model"
)

type JsonParser struct {
}

func (p *JsonParser) Transfer(content string) (*model.Field, error) {
	data := make(map[string]interface{}, 0)
	err := json.Unmarshal([]byte(content), &data)
	if err != nil {
		return nil, err
	}

	root := &model.Field{FieldType: "root", ChildField: make([]*model.Field, 0)}

	for key, value := range data {
		switch value.(type) {
		case map[string]interface{}:
			dealMap(root, key, value.(map[string]interface{}))
		case []interface{}:
			dealArray(root, key, value.([]interface{}))
		default:
			field := &model.Field{
				FieldName: key,
				FieldType: value.(string),
			}
			root.ChildField = append(root.ChildField, field)
		}

	}

	return root, nil
}

// 循环解析map结构
func dealMap(root *model.Field, key string, items map[string]interface{}) {
	field := &model.Field{
		FieldName:  key,
		FieldType:  constants.Object,
		ChildField: make([]*model.Field, 0, len(items)),
	}
	for k, v := range items {
		switch v.(type) {
		case []interface{}:
			dealArray(field, k, v.([]interface{}))
		default:
			field.ChildField = append(field.ChildField, &model.Field{
				FieldName: k,
				FieldType: v.(string),
			})
		}
	}
	root.ChildField = append(root.ChildField, field)
}

// 循环解析array
func dealArray(root *model.Field, key string, items []interface{}) {
	field := &model.Field{
		FieldName:  key,
		FieldType:  constants.Array,
		ChildField: make([]*model.Field, 0, len(items)),
	}

	for _, item := range items {
		switch item.(type) {
		case map[string]interface{}:
			dealMap(field, constants.None, item.(map[string]interface{}))
		case string:
			field.ChildField = append(field.ChildField, &model.Field{
				FieldType: item.(string),
			})
		}
	}

	root.ChildField = append(root.ChildField, field)
}
