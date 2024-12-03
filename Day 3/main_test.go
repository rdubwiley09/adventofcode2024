package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadText(t *testing.T) {
	assert := assert.New(t)

	command := LoadText("./test.txt")
	assert.Equal(command, "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))")
}

func TestGetMulSections(t *testing.T) {
	assert := assert.New(t)

	command := LoadText("./test.txt")
	mulSections := GetMulSections(command)
	assert.Equal(len(mulSections), 4)
}
