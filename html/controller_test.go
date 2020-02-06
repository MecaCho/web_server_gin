package html

import (
	"fmt"
	"testing"
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
