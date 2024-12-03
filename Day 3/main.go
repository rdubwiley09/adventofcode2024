package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func LoadText(loc string) string {
	file, err := os.ReadFile(loc)
	if err != nil {
		log.Fatal(err)
	}
	return string(file)
}

func GetMulSections(command string) []string {
	r, _ := regexp.Compile("mul\\([0-9]{1,3},[0-9]{1,3}\\)")
	return r.FindAllString(command, -1)
}

func RunMult(subCommand string) int {
	output := 1
	r, _ := regexp.Compile("[0-9]{1,3}")
	nums := r.FindAllString(subCommand, 2)
	for _, item := range nums {
		newInt, _ := strconv.Atoi(item)
		output *= newInt
	}
	return output
}

func CalculateMult(subCommands []string) int {
	output := 0
	for _, item := range subCommands {
		output += RunMult(item)
	}
	return output
}

func RunCommandProcess(command string) int {
	subCommands := GetMulSections(command)
	return CalculateMult(subCommands)
}

func CorrectCommand(command string) []string {
	r, _ := regexp.Compile("mul\\(\\d+,\\d+\\)|do\\(\\)|don't\\(\\)")
	commands := r.FindAllString(command, -1)
	output := []string{}
	addCommand := true
	for _, command := range commands {
		if command == "don't()" {
			addCommand = false
		} else if command == "do()" {
			addCommand = true
		} else {
			if addCommand {
				output = append(output, command)
			}
		}
	}
	return output
}

func main() {
	command := LoadText("./data.txt")
	fmt.Println(RunCommandProcess(command))
	correctedCommand := CorrectCommand(command)
	fmt.Println(CalculateMult(correctedCommand))
}
