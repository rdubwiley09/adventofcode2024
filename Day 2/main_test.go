package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadLists(t *testing.T) {
	assert := assert.New(t)

	report := LoadLists("./test.txt")
	assert.Equal(len(report[0]) > 0, true)
}

func TestGetSublists(t *testing.T) {
	assert := assert.New(t)
	report := [][]int{{1, 2, 3}}
	subLists := GetSubLists(report[0])
	assert.Equal(len(subLists), 3)
}
