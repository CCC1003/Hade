package session

import (
	"Hade/Horm/dialect"
	"database/sql"
	"os"
	"testing"
)

var (
	TestDB      *sql.DB
	TestDial, _ = dialect.GetDialect("mysql")
)

func TestMain(m *testing.M) {
	TestDB, _ = sql.Open("mysql", "root:123456@tcp(8.130.85.112:3306)/Horm")
	code := m.Run()
	_ = TestDB.Close()
	os.Exit(code)
}

func NewSession() *Session {
	return New(TestDB, TestDial)
}
