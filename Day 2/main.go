package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func LoadLists(loc string) [][]int {
	output := [][]int{}
	file, err := os.Open(loc)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		result := strings.Split(line, " ")
		newSlice := []int{}
		for _, num := range result {
			newInt, _ := strconv.Atoi(num)
			newSlice = append(newSlice, newInt)
		}
		output = append(output, newSlice)
	}

	if scanerr := scanner.Err(); scanerr != nil {
		log.Fatal(scanerr)
	}
	return output
}

func CalculateDifference(numList []int) []int {
	output := []int{}
	for i, item := range numList[:len(numList)-1] {
		output = append(output, item-numList[i+1])
	}
	return output
}

func IsIncreasingOrDecreasing(differencelist []int) int {
	minDifference := slices.Min(differencelist)
	maxDifference := slices.Max(differencelist)
	if slices.Contains(differencelist, 0) {
		return 0
	}
	if minDifference < 0 && maxDifference > 0 {
		return 0
	}
	return 1
}

func IsWithinDifferenceLimit(differnceList []int) int {
	for _, item := range differnceList {
		if item < -3 {
			return 0
		}
		if item > 3 {
			return 0
		}
	}
	return 1
}

func RunItemSafety(report []int) int {
	differenceList := CalculateDifference(report)
	incDecCondition := IsIncreasingOrDecreasing(differenceList)
	differenceCondition := IsWithinDifferenceLimit(differenceList)
	return incDecCondition * differenceCondition
}

func RunSafeCount(reportList [][]int) int {
	count := 0
	for _, report := range reportList {
		count += RunItemSafety(report)
	}
	return count
}

func GetSubLists(numList []int) [][]int {
	outputList := [][]int{}
	for i, _ := range numList {
		destination := make([]int, len(numList))
		copy(destination, numList)
		destination = slices.Delete(destination, i, i+1)
		outputList = append(outputList, destination)
	}
	return outputList
}

func RunDamperCount(reportList [][]int) int {
	count := 0
	for _, report := range reportList {
		safeCount := 0
		safeCount += RunItemSafety(report)
		subLists := GetSubLists(report)
		safeCount += RunSafeCount(subLists)
		if safeCount > 0 {
			count += 1
		}
	}
	return count
}

func main() {
	reports := LoadLists("./data.txt")
	safeCount := RunSafeCount(reports)
	fmt.Println(safeCount)
	damperCount := RunDamperCount(reports)
	fmt.Println(damperCount)
}
