package main

import (
	"database/sql"
	"github.com/astaxie/beego/toolbox"
)

type DatabaseCheck struct {
}

func (*DatabaseCheck) Check() error {
	_, err := sql.Open("mysql", "web:cx6222580@tcp(139.198.181.33:13306)/test")
	if err != nil {
		return err
	}
	return nil
}

func init() {
	toolbox.AddHealthCheck("database", &DatabaseCheck{})
}
