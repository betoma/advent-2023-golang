package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

func loadCards(filename string) (cards [][2][]string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		sep := strings.Split(scanner.Text(), ": ")
		cardsep := strings.Split(sep[1], " | ")
		winning_nums := strings.Fields(cardsep[0])
		my_nums := strings.Fields(cardsep[1])
		cards = append(cards, [2][]string{winning_nums, my_nums})
	}
	return
}

func partOne(cards [][2][]string) (total int) {
	for _, c := range cards {
		var score int
		for _, n := range c[1] {
			if slices.Contains(c[0], n) {
				if score < 1 {
					score = 1
				} else {
					score *= 2
				}
			}
		}
		total += score
	}
	return
}

func partTwo(cards [][2][]string) (total int) {
	copies := make(map[int]int)
	for n := 0; n < len(cards); n++ {
		copies[n] = 1
	}
	for i, c := range cards {
		var n_matches int
		for _, n := range c[1] {
			if slices.Contains(c[0], n) {
				n_matches += 1
			}
		}
		for n := 1; n <= n_matches; n++ {
			copies[i+n] += 1 * copies[i]
		}
	}
	for _, quant := range copies {
		total += quant
	}
	return
}

func main() {
	cards := loadCards("input.txt")
	p1 := partOne(cards)
	fmt.Println("Part One: ", p1)
	p2 := partTwo(cards)
	fmt.Println("Part Two: ", p2)
}
