package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type list struct {
	temp []string
}

type data struct {
	ReactionTitle string `json:"reaction_title"`
	VoltTitle     string `json:"volt_title"`
	reaction      list   `json:"reactions"`
	voltage       list   `json:"volts"`
}

func initData() (data, error) {
	//dataFilepath := "data/seed.json"
	file, err := os.ReadFile("data/seed.json")
	//fmt.Println(err)

	var electrochemicals data

	err = json.Unmarshal(file, &electrochemicals)
	//fmt.Println(err)
	fmt.Println(electrochemicals)

	return electrochemicals, err
}

func search(s string) (bool, int, string, string) {
	memory, _ := initData()

	reactionStack := memory.reaction
	voltStack := memory.voltage
	fmt.Println(reactionStack)

	low := 0
	high := len(reactionStack.temp) - 1

	for low <= high {
		median := (low + high) / 2

		if reactionStack.temp[median] < s {
			low = median + 1
		} else {
			high = median - 1
		}
	}

	if low == len(reactionStack.temp) || reactionStack.temp[low] != s {
		return false, -1, "", ""
	}

	index := getIndex(s, reactionStack.temp)

	return true, index, reactionStack.temp[index], voltStack.temp[index]

}

func getIndex(element string, reactionStack []string) int {
	index := 0
	for _, reaction := range reactionStack {
		index = strings.Index(reaction, element)
	}
	return index
}

func main() {

	arg := os.Args[1]

	isPresent, _, reaction, volt := search(arg)
	fmt.Println(isPresent)

	if isPresent == false {
		fmt.Println("Element not a part of the Electrochemical series")
	} else {
		fmt.Printf("Element %s found \n Reaction: %s \n Volts: %s\n", arg, reaction, volt)
	}

}
