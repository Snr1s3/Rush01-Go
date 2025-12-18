package main

import (
	"fmt"
	"os"
)

var Dict map[string]string
var Str string = " "
var Error bool = false

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run rush02.go <cluestring> <language>")
		return
	}
	if ReadFile(os.Args[2]) {
		return
	}
	processNumber(os.Args[1])
}

func processNumber(numbers string) {
	cleaned := removeLeadingZeros(numbers)
	if cleaned == "" {
		search("0")
		fmt.Print(Str, "\n")
		return
	}
	if len(cleaned) == 1 {
		search(string(cleaned[0]))
		fmt.Print(Str, "\n")
		return
	}
	printNumberGroups(cleaned)
	fmt.Print(Str, "\n")
}

func removeLeadingZeros(numbers string) string {
	i := 0
	for i < len(numbers) && numbers[i] == '0' {
		i++
	}
	return numbers[i:]
}

func printNumberGroups(numbers string) {
	leng := len(numbers)
	len2 := leng
	for n := 0; n < leng; n++ {
		len2 = len2 - n
		mod := len2 % 3
		str := string(numbers[n])
		if str == "0" {
			continue
		} else if mod == 1 && n != leng-1 && n != leng-2 {
			errorControl(str, 0)
			errorControl("1", leng-n)
		} else if mod == 1 && n != leng-1 {
			errorControl(str, leng-n)
		} else if mod == 0 {
			errorControl(str, 0)
			errorControl("1", 3)
		} else if mod == 2 {
			errorControl(str, 2)
		} else if n == leng-2 {
			if str == "1" {
				str += string(numbers[n+1])
				errorControl(str, 0)
				break
			}
			errorControl(str, leng-n)
		} else {
			errorControl(str, 0)
		}
		if Error {
			return
		}
	}
}

func errorControl(str string, leng int) {
	if leng == 0 {
		search(str)
	} else {
		search(addZeros(str, leng))
	}
}

func search(key string) {
	value, exists := Dict[key]
	if exists {
		Str += value
	} else {
		fmt.Print("Dict Error\n")
		Error = true
	}
}

func addZeros(start string, leng int) string {
	str := start
	for n := 0; n < leng-1; n++ {
		str += "0"
	}
	return str
}
