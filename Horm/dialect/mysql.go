package dialect

import "reflect"

type mysql struct{}

var _ Dialect = (*mysql)(nil)

func init() {
	RegisterDialect("mysql", &mysql{})
}

func (m *mysql) DataTypeOf(typ reflect.Value) string {

	return ""
}
func (m *mysql) TableExistSQL(tableName string) (string, []interface{}) {
	return "", nil
}

