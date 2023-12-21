package schema

import (
	"Hade/Horm/dialect"
	"testing"
)

type User struct {
	Name string `horm:"primary key"`
	Age  int
}

var TestDial, _ = dialect.GetDialect("mysql")

func TestParse(t *testing.T) {
	schema := Parse(&User{}, TestDial)
	if schema.Name != "User" || len(schema.Fields) != 2 {
		t.Fatal("failed to parse User struct")
	}
	if schema.GetField("Name").Tag != "primary key" {
		t.Fatal("failed to parse primary key")
	}
}
