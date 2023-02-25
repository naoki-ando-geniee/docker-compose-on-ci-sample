package main

import (
	"database/sql"
	"fmt"
	"testing"
)

func TestMysql(t *testing.T) {
	client, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/test")
	if err != nil {
		t.Fatal(err)
	}
	err = client.Ping()
	if err != nil {
		t.Fatal(err)
	}
	row := client.QueryRow("SELECT 'hoge'")
	var hoge string
	err = row.Scan(&hoge)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(hoge)
}
