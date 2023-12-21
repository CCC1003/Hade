package session

import (
	"Hade/Horm/log"
	"Hade/Horm/schema"
	"fmt"
	"reflect"
	"strings"
)

func (s *Session) Model(value interface{}) *Session {
	//nil or different model,update refTable
	if s.refTable == nil || reflect.TypeOf(value) != reflect.TypeOf(s.refTable.Model) {
		s.refTable = schema.Parse(value, s.dialect)
	}
	return s
}
func (s *Session) RefTable() *schema.Schema {
	if s.refTable == nil {
		log.Error("Model is not set")
	}
	return s.refTable
}

// CreateTable create table
func (s *Session) CreateTable() error {
	table := s.RefTable()
	var columns []string
	for _, field := range table.Fields {
		columns = append(columns, fmt.Sprintf("%s %s %s", field.Name, field.Type, field.Tag))
	}
	desc := strings.Join(columns, ",")
	_, err := s.Raw(fmt.Sprintf("CREATE TABLE %s (%s);", s.RefTable().Name, desc)).Exec()
	return err
}

// DropTable drop table
func (s *Session) DropTable() error {
	_, err := s.Raw(fmt.Sprintf("DROP TABLE IF EXISTS %s", s.RefTable().Name)).Exec()
	return err
}
func (s *Session) HasTable() bool {
	var dbName string
	err := s.db.QueryRow("SELECT DATABASE()").Scan(&dbName)
	if err != nil {
		fmt.Println("Failed to get MySQL database name:", err)
		return false
	}

	sql, value := s.dialect.TableExistSQL(dbName, s.refTable.Name)
	row := s.Raw(sql, value...).QueryRow()
	var tmp string
	_ = row.Scan(&tmp)

	return tmp == s.RefTable().Name
}
