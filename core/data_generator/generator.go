package data_generator

import (
	"mongo_data/constants"
	"mongo_data/model"
)

func GenerateData(data map[string]interface{}, root *model.Field, generator Generator) {
	var res interface{}
	var tmp, newData map[string]interface{}
	for _, item := range root.ChildField {
		switch item.FieldType {
		case constants.Int32:
			res = generator.GetInt32()
		case constants.Int64:
			res = generator.GetInt64()
		case constants.Bool:
			res = generator.GetBool()
		case constants.String:
			res = generator.GetString()
		case constants.Double:
			res = generator.GetDouble()
		case constants.Date:
			res = generator.GetDate()
		case constants.Array:
			data[item.FieldName] = make([]interface{}, 0)
		case constants.Object:
			newData = make(map[string]interface{}, 0)
			if item.FieldName == constants.None {
				data[root.FieldName] = append(data[root.FieldName].([]interface{}), newData)
			} else {
				data[item.FieldName] = newData
			}

			tmp = data
			data = newData

		}

		if root.FieldType == constants.Array && isBasicType(item.FieldType) {
			data[root.FieldName] = append(data[root.FieldName].([]interface{}), res)
		}

		if root.FieldType == constants.Object && isBasicType(item.FieldType) {
			data[item.FieldName] = res
		}

		if isBasicType(root.FieldType) && isBasicType(item.FieldType) {
			data[item.FieldName] = res
		}

		GenerateData(data, item, generator)

		// 递归回来之后需要将data恢复为递归前的状态
		if tmp != nil {
			data = tmp
			tmp = nil
		}

	}
}

func isBasicType(t string) bool {
	if t != constants.Object && t != constants.Array {
		return true
	}
	return false
}
