package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func getAdjacentCoords(line_no int, start int, end int) (next [][2]int) {
	next = append(next, [2]int{start - 1, line_no}, [2]int{end, line_no})
	for i := start - 1; i <= end; i++ {
		next = append(next, [2]int{i, line_no - 1}, [2]int{i, line_no + 1})
	}
	return
}

func parseSchematic(filename string) (number_locs map[[3]int]int, symbol_locs map[[2]int]string) {
	number_locs = make(map[[3]int]int)
	symbol_locs = make(map[[2]int]string)
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	j := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		currently_digit := false
		var dig_start int
		for i, char := range strings.Split(scanner.Text(), "") {
			if _, err := strconv.Atoi(char); err == nil {
				if !currently_digit {
					dig_start = i
					currently_digit = true
				}
			} else {
				if currently_digit {
					number, err := strconv.Atoi(scanner.Text()[dig_start:i])
					if err != nil {
						log.Fatal(err)
					}
					number_locs[[3]int{j, dig_start, i}] = number
					currently_digit = false
				}
				if char != "." && char != "\n" {
					symbol_locs[[2]int{i, j}] = char
				}
			}
		}
		if currently_digit {
			number, err := strconv.Atoi(scanner.Text()[dig_start:])
			if err != nil {
				log.Fatal(err)
			}
			number_locs[[3]int{j, dig_start, len(scanner.Text())}] = number
		}
		j++
	}

	return
}

func partOne(numbers map[[3]int]int, symbols map[[2]int]string) (sum int) {
	sum = 0
	for loc, num := range numbers {
		adjacent := getAdjacentCoords(loc[0], loc[1], loc[2])
		// log.Println("Number: ", num, " at ", "[", loc[1], ":", loc[2], ", ", loc[0], "], adjacent cells: ", adjacent)
		for _, coord := range adjacent {
			if _, ok := symbols[coord]; ok {
				// log.Println("Added!")
				sum += num
				break
			}
		}
	}
	return
}

func partTwo(numbers map[[3]int]int, symbols map[[2]int]string) (sum int) {
	sum = 0
	asterisks := make(map[[2]int][]int)
	for loc, num := range numbers {
		adjacent := getAdjacentCoords(loc[0], loc[1], loc[2])
		for _, coord := range adjacent {
			if sym, ok := symbols[coord]; ok {
				if sym == "*" {
					asterisks[coord] = append(asterisks[coord], num)
				}
			}
		}
	}
	for _, numbers := range asterisks {
		if len(numbers) == 2 {
			sum += numbers[0] * numbers[1]
		}
	}
	return
}

func main() {
	numbers, symbols := parseSchematic("input.txt")
	p1 := partOne(numbers, symbols)
	fmt.Println("Part One: ", p1)
	p2 := partTwo(numbers, symbols)
	fmt.Println("Part Two: ", p2)
}
