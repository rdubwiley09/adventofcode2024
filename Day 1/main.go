package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func LoadLists(loc string) ([]int, []int) {
	firstList := []int{}
	secondList := []int{}
	file, err := os.Open(loc)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		result := strings.SplitN(line, "   ", 2)
		int1, ferr := strconv.Atoi(result[0])
		if ferr != nil {
			log.Fatal(ferr)
		}
		int2, serr := strconv.Atoi(result[1])
		if serr != nil {
			log.Fatal(serr)
		}
		firstList = append(firstList, int1)
		secondList = append(secondList, int2)
	}

	if scanerr := scanner.Err(); scanerr != nil {
		log.Fatal(scanerr)
	}
	return firstList, secondList
}

func SortList(rawList []int) []int {
	sort.Ints(rawList)
	return rawList
}

func CompareLists(firstList []int, secondList []int) int {
	difference := 0
	itemDif := 0
	for i, num := range firstList {
		itemDif = num - secondList[i]
		if itemDif < 0 {
			difference -= itemDif
		} else {
			difference += itemDif
		}
	}
	return difference
}

func RunComparison(firstList []int, secondList []int) int {
	firstSorted := SortList(firstList)
	secondSorted := SortList(secondList)
	return CompareLists(firstSorted, secondSorted)
}

func MakeMap(secondList []int) map[int]int {
	dict := make(map[int]int)
	for _, num := range secondList {
		// dict[num] will return 0 if it hasn't yet been initialised
		dict[num] = dict[num] + 1
	}
	return dict
}

func ComputeSimilarity(firstList []int, dict map[int]int) int {
	similarity := 0
	for _, num := range firstList {
		similarity += num * dict[num]
	}
	return similarity
}

func RunSimilarity(firstList []int, secondList []int) int {
	dict := MakeMap(secondList)
	return ComputeSimilarity(firstList, dict)
}

func main() {
	firstList, secondList := LoadLists("./data.txt")
	comparison := RunComparison(firstList, secondList)
	similarity := RunSimilarity(firstList, secondList)
	fmt.Println("Comparison is:")
	fmt.Println(comparison)
	fmt.Println("Simlilarity is:")
	fmt.Println(similarity)
}
