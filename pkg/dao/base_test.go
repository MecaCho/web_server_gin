package dao

import (
	"fmt"
	"testing"
)

func TestDB_FilterTable(t *testing.T) {
	filters := map[string][]string{}
	// filters[]
	if filters != nil {
		fmt.Println(filters, "not nil")
	} else {
		fmt.Println(filters, "nil")
	}
}
