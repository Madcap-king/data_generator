package model

type Field struct {
	FieldName  string
	FieldType  string
	ChildField []*Field
}
