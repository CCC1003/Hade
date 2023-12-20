package schema

import (
	"Hade/Horm/dialect"
)

type Field struct {
	Name string
	Type string
	Tag  string
}

type Schema struct {
	Model      interface{}
	Name       string
	Fields     []*Field
	FieldNames []string
	fieldMap   map[string]*Field
}

func (schema *Schema) GetField(name string) *Field {
	return schema.fieldMap[name]
}
func Parse(dest interface{}, d dialect.Dialect) {
	//modelType := reflect.Indirect(reflect.ValueOf(dest)).Type()

}
