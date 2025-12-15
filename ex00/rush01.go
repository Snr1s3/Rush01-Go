package main

import (
	"fmt"
	"os"
	"strconv"
)

var Map [][]int
var Rows int = 4
var Cols int = 4

func main() {
	input := os.Args[1:]
	if len(input) != Rows*Cols {
		fmt.Println("Invalid number of arguments")
		return
	}
	fmt.Println("Allocating")
	Map = make([][]int, Rows)
	for i := range Map {
		Map[i] = make([]int, Cols)
	}
	count := 0
	for r := 0; r < Rows; r++ {
		for c := 0; c < Cols; c++ {
			val, err := strconv.Atoi(input[count])
			if err != nil {
				fmt.Printf("Invalid integer: %s\n", input[count])
				return
			}
			Map[r][c] = val
			count++
		}
	}
	PrintMap()
}

func PrintMap() {
	for c := 0; c < Cols; c++ {
		fmt.Print(" ")
		fmt.Print(Map[0][c])
	}
	fmt.Print("\n")
	for c := 0; c < Cols; c++ {
		fmt.Print(Map[2][c])
		fmt.Print("       ")
		fmt.Print(Map[3][c])
		fmt.Print("\n")
	}
	for c := 0; c < Cols; c++ {
		fmt.Print(" ")
		fmt.Print(Map[1][c])
	}
	fmt.Print("\n")
}
