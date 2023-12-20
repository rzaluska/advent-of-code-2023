package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type SearchTree struct {
	children       []*SearchTree
	startCharacter byte
}

func (root *SearchTree) Add(s string) {
	if s == "" {
		return
	}
	for _, c := range root.children {
		if c.startCharacter == s[0] {
			c.Add(s[1:])
			return
		}
	}
	newChild := &SearchTree{}
	newChild.startCharacter = s[0]
	root.children = append(root.children, newChild)
	newChild.Add(s[1:])
}

func (root *SearchTree) isLeaf() bool {
	return len(root.children) == 0
}

func (root *SearchTree) MatchPrefix(s string) (string, bool) {
	if s == "" {
		//we processed everything and we return true if we got to leaf node
		return "", root.isLeaf()
	}
	for _, c := range root.children {
		if c.startCharacter == s[0] {
			r, b := c.MatchPrefix(s[1:])
			return string(c.startCharacter) + r, b
		}
	}
	//we can't match any of child nodes
	//if we are leaf that means our match is complete from root to leaf
	return "", root.isLeaf()
}

func NumStringToInt(s string) string {
	switch s {
	case "one":
		return "1"
	case "two":
		return "2"
	case "three":
		return "3"
	case "four":
		return "4"
	case "five":
		return "5"
	case "six":
		return "6"
	case "seven":
		return "7"
	case "eight":
		return "8"
	case "nine":
		return "9"
	}
	return ""
}

func PrepareSearchTree() *SearchTree {
	st := &SearchTree{}
	values := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	for _, v := range values {
		st.Add(v)
	}
	return st

}

func PreprocessLine(st *SearchTree, input string) string {
	inputCopy := strings.Clone(input)
	var output strings.Builder
	var wordLen int
	for len(inputCopy) > 0 {
		match, isComplete := st.MatchPrefix(inputCopy)
		if isComplete {
			output.WriteString(NumStringToInt(match))
			inputCopy = inputCopy[1:]
			wordLen = len(match) - 1
		} else if !isComplete && wordLen > 0 {
			inputCopy = inputCopy[1:]
			wordLen--
		} else if !isComplete {
			output.WriteByte(inputCopy[0])
			inputCopy = inputCopy[1:]
		}
	}
	return output.String()

}

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

var fileName = flag.String("f", "", "")
var preprocess = flag.Bool("p", false, "")

func main() {
	flag.Parse()
	if *fileName == "" {
		flag.Usage()
		return
	}
	f, err := os.Open(*fileName)
	if err != nil {
		log.Fatalf("Error opening file %s: %s", *fileName, err)
	}
	scanner := bufio.NewScanner(f)
	st := PrepareSearchTree()
	var sum int
	for scanner.Scan() {
		line := scanner.Text()
		if *preprocess {
			line = PreprocessLine(st, line)
			fmt.Printf("%s\n", line)
		}
		sum += processLine(line)
	}
	fmt.Printf("Sum: %d\n", sum)
}
