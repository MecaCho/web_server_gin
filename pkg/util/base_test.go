package util

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenPwd(t *testing.T) {
	ret := GenPwd(16)
	fmt.Println(ret)
	assert.Equal(t, 16, len(ret), "Not match")
}
