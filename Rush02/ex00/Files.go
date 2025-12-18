package main

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFile(language string) bool {
	path := "lan/numbers." + language + ".dict"
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return true
	}
	defer file.Close()

	Dict = make(map[string]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := splitByColon(line)
		if len(parts) == 2 {
			key := parts[0]
			value := parts[1]
			Dict[key] = value
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return true
	}
	return false
}

func splitByColon(s string) []string {
	for i := 0; i < len(s); i++ {
		if s[i] == ':' {
			return []string{
				s[:i],
				s[i+1:],
			}
		}
	}
	return []string{s}
}
