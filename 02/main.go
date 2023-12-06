package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"golang.org/x/exp/maps"
)

func cubeCount(filename string) (total int, power int) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	color_counts := map[string]int{"red": 12, "green": 13, "blue": 14}
	total = 0
	power = 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "Game") {
			invalid := false
			min_cubes := make(map[string]int)
			first_split := strings.Split(scanner.Text(), ":")
			game_no, err := strconv.Atoi(strings.Fields(first_split[0])[1])
			if err != nil {
				log.Fatal(err)
			}
			games_list := strings.Split(first_split[1], ";")
			for _, game := range games_list {
				cubes_list := strings.Split(game, ",")
				for _, cubeset := range cubes_list {
					cubesplit := strings.Fields(cubeset)
					n, err := strconv.Atoi(cubesplit[0])
					if err != nil {
						log.Fatal(err)
					}
					color := cubesplit[1]
					if current_min, found := min_cubes[color]; found {
						if n > current_min {
							min_cubes[color] = n
						}
					} else {
						min_cubes[color] = n
					}
					if max_val, found := color_counts[color]; found {
						if n > max_val {
							// log.Println("Number of ", color, " beads (", n, ") is greater than ", max_val)
							invalid = true
						}
					} else {
						// log.Println("Color not in dict")
						invalid = true
					}
				}
			}
			if !invalid {
				total += game_no
			}
			p := 1
			for _, n := range maps.Values(min_cubes) {
				p *= n
			}
			power += p
		}
	}
	return
}

func main() {
	p1, p2 := cubeCount("input.txt")
	fmt.Println("Part One: ", p1)
	fmt.Println("Part Two: ", p2)
}
