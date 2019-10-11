package db

import "testing"

func TestDbConnection_ConnectionURL(t *testing.T) {
	conn := &Connection{
		DbName:     "lab2",
		User:       "lab2user",
		Password:   "test",
		Host:       "localhost",
		DisableSSL: true,
	}
	if conn.ConnectionURL() != "postgres://lab2user:test@localhost/lab2?sslmode=disable" {
		t.Error("Unexpected connection string")
	}
}
