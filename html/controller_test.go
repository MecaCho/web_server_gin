package html

import (
	"fmt"
	"testing"
	"time"
)

func TestServerHandle_IndexController(t *testing.T) {
	var testStr string
	testStr = "123fggjg"
	if len(testStr) > 10 {
		fmt.Println(testStr[:10])
	} else {
		fmt.Println(testStr)
	}
}

func TestServerHandle_IndexController2(t *testing.T) {
	now := time.Now()
	nowStr := now.Format("2006-01")
	fmt.Println(nowStr)
}
