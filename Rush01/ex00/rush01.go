package main

import (
	"fmt"
	"os"
	"strconv"
)

var Views [][]int
var Map [][]int
var NUM int = 4

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run rush01.go <cluestring>")
		return
	}
	clues := os.Args[1]
	if len(clues) != 16 && len(clues) != 20 && len(clues) != 24 {
		fmt.Println("Clue string must be 16/20/24 digits")
		return
	}
	var input []string
	if len(clues) == 16 {
		input = make([]string, 16)
		NUM = 4
	} else if len(clues) == 20 {
		input = make([]string, 20)
		NUM = 5
	} else if len(clues) == 24 {
		input = make([]string, 24)
		NUM = 6
	} else {
		fmt.Println("Clue string must be 16/20/24 digits")
		return
	}
	for i := 0; i < NUM*4; i++ {
		input[i] = string(clues[i])
	}
	if !Allocate(input) {
		return
	}
	fmt.Println()
	fmt.Println("Map to solve:")
	PrintMap()
	if !Solve(0, 0) {
		fmt.Println("No solution found for the given clues.")
	}
	fmt.Println()
	fmt.Println("Map solved:")
	PrintMap()
}
func Allocate(input []string) bool {
	fmt.Println("Allocating Arrays in Memory")
	Map = make([][]int, NUM)
	for i := range Map {
		Map[i] = make([]int, NUM)
		for j := range Map[i] {
			Map[i][j] = 0
		}
	}
	Views = make([][]int, 4)
	for i := range Views {
		Views[i] = make([]int, NUM)
	}
	count := 0
	for i := 0; i < 4; i++ {
		for j := 0; j < NUM; j++ {
			val, err := strconv.Atoi(input[count])
			if err != nil || val > NUM || val < 1 {
				fmt.Printf("Invalid integer: %s\n", input[count])
				return false
			}
			Views[i][j] = val
			count++
		}
	}
	return true
}

func PrintMap() {
	PrintTopBotViews(0)
	PrintTopBotLine()
	PrintMiddle()
	PrintTopBotViews(1)
}
func PrintMiddle() {
	for r := 0; r < NUM; r++ {
		fmt.Printf(" %d |", Views[2][r])
		for c := 0; c < NUM; c++ {
			if Map[r][c] == 0 {
				fmt.Print("   |")
			} else {
				fmt.Printf(" %d |", Map[r][c])
			}
		}
		fmt.Printf(" %d\n", Views[3][r])
		PrintTopBotLine()
	}
}
func PrintTopBotLine() {
	fmt.Print("   +")
	for c := 0; c < NUM; c++ {
		fmt.Print("---+")
	}
	fmt.Println()
}
func PrintTopBotViews(view int) {
	fmt.Print("     ")
	for c := 0; c < NUM; c++ {
		fmt.Printf("%d   ", Views[view][c])
	}
	fmt.Println()
}
