package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadLists(t *testing.T) {
	assert := assert.New(t)

	firstList, secondList := LoadLists("./test.txt")
	assert.Equal(firstList, []int{3, 4, 2, 1, 3, 3})
	assert.Equal(secondList, []int{4, 3, 5, 3, 9, 3})
}

func TestSortLists(t *testing.T) {
	assert := assert.New(t)
	firstList, secondList := LoadLists("./test.txt")
	firstSorted := SortList(firstList)
	secondSorted := SortList(secondList)
	assert.Equal(firstSorted, []int{1, 2, 3, 3, 3, 4})
	assert.Equal(secondSorted, []int{3, 3, 3, 4, 5, 9})
}

func TestCompareLists(t *testing.T) {
	assert := assert.New(t)
	firstList, secondList := LoadLists("./test.txt")
	comparison := RunComparison(firstList, secondList)
	assert.Equal(comparison, 11)
}

func TestMakeMap(t *testing.T) {
	assert := assert.New(t)
	_, secondList := LoadLists("./test.txt")
	dict := MakeMap(secondList)
	assert.Equal(dict[3], 3)
	t.Log(dict[0])
}

func TestRunSimilarity(t *testing.T) {
	assert := assert.New(t)
	firstList, secondList := LoadLists("./test.txt")
	similarity := RunSimilarity(firstList, secondList)
	assert.Equal(similarity, 31)
}
