package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func processLine(line string) int {
	var first rune
	var last rune
	for _, char := range line {
		if unicode.IsDigit(char) {
			first = char
			break
		}
	}
	for i := range line {
		char := rune(line[len(line)-i-1])
		if unicode.IsDigit(char) {
			last = char
			break
		}
	}
	result, err := strconv.Atoi(string(first) + string(last))
	if err != nil {
		log.Fatalf("error converting from string %s to int %s", line, err)
	}
	return result
}

func main() {
	f, err := os.Open("day1.txt")
	if err != nil {
		log.Fatalf("Error opening file %s: %s", "day1.txt", err)
	}
	scanner := bufio.NewScanner(f)
	var sum int
	for scanner.Scan() {
		sum += processLine(scanner.Text())
	}
	fmt.Printf("Sum: %d\n", sum)
}
