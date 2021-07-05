package service

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCode(t *testing.T) {
	a, err := GetCode("abcdefghijklmnop")
	assert.Nil(t, err)
	fmt.Println(a)
}

func TestJwtEncrypt(t *testing.T) {

}
