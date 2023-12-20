package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	MaxRed   int = 12
	MaxGreen int = 13
	MaxBlue  int = 14
)

type Subset struct {
	Red   int
	Green int
	Blue  int
}

type Game struct {
	Id      int
	Subsets []Subset
}

func NewGame(id int) *Game {
	var object Game
	object.Id = id
	object.Subsets = make([]Subset, 0)
	return &object
}

func parseGameId(input string) (result int) {
	parts := strings.Split(input, " ")
	if len(parts) != 2 {
		log.Fatalf("error parsing GameId %s, space separator not found", input)
	}
	result, err := strconv.Atoi(parts[1])
	if err != nil {
		log.Fatalf("error parsing GameId %s, space separator not found", input)
	}
	return result
}

func parseSubsets(input string) []Subset {
	parts := strings.Split(input, "; ")
	result := make([]Subset, len(parts))
	for i, part := range parts {
		result[i] = parseSubset(part)
	}
	return result
}

func parseSubset(input string) (s Subset) {
	parts := strings.Split(input, ", ")
	for _, part := range parts {
		nameAndNumber := strings.Split(part, " ")
		number, err := strconv.Atoi(nameAndNumber[0])
		if err != nil {
			log.Fatalf("cant parse subset %s", part)
		}
		name := nameAndNumber[1]
		switch name {
		case "red":
			s.Red = number
		case "green":
			s.Green = number
		case "blue":
			s.Blue = number
		}
	}
	return
}

func parseGame(input string) *Game {
	parts := strings.Split(input, ": ")
	if len(parts) != 2 {
		log.Fatalf("error parsing line %s, : separator not found", input)
	}
	gameId := parseGameId(parts[0])
	game := NewGame(gameId)
	game.Subsets = parseSubsets(parts[1])
	return game
}

func gamePossible(g *Game) bool {
	for _, subset := range g.Subsets {
		if subset.Red > MaxRed {
			return false
		}
		if subset.Blue > MaxBlue {
			return false
		}
		if subset.Green > MaxGreen {
			return false
		}
	}
	return true
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("error opening file %s: %s", "input.txt", err)
	}
	scanner := bufio.NewScanner(file)
	var sum int
	for scanner.Scan() {
		line := scanner.Text()
		game := parseGame(line)
		if gamePossible(game) {
			fmt.Printf("%+v\n", game)
			sum += game.Id
		}
		fmt.Printf("%d\n", sum)
	}
}
