// First day of advent of code 2023 solution
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func partOne(filename string) (total int) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	total = 0
	for scanner.Scan() {
		li := []rune(scanner.Text())
		var first_digit rune
		var second_digit rune
		for _, c := range li {
			if unicode.IsDigit(c) {
				first_digit = c
				break
			}
		}
		for index := len(li) - 1; index >= 0; index-- {
			r := li[index]
			if unicode.IsDigit(r) {
				second_digit = r
				break
			}
		}
		n, err := strconv.Atoi(string([]rune{first_digit, second_digit}))
		if err != nil {
			log.Fatal(err)
		}
		total += n
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return
}

func partTwo(filename string) (total int) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	starters := map[rune]any{
		't': map[rune]rune{'w': '\u0032', 'h': '\u0033'},
		'f': map[rune]rune{'o': '\u0034', 'i': '\u0035'},
		's': map[rune]rune{'i': '\u0036', 'e': '\u0037'},
		'o': '\u0031',
		'e': '\u0038',
		'n': '\u0039',
	}
	spelling := map[rune][]rune{
		'\u0031': []rune("one"),
		'\u0032': []rune("two"),
		'\u0033': []rune("three"),
		'\u0034': []rune("four"),
		'\u0035': []rune("five"),
		'\u0036': []rune("six"),
		'\u0037': []rune("seven"),
		'\u0038': []rune("eight"),
		'\u0039': []rune("nine"),
	}
	scanner := bufio.NewScanner(file)
	total = 0
	for scanner.Scan() {
		// log.Println("current line: ", scanner.Text())
		first_digit_found := false
		second_digit_found := false
		var first_digit rune
		var second_digit rune
		word_idx := []int{}
		word_cat := []rune{}
		for _, c := range scanner.Text() {
			// log.Println("char #", j)
			if unicode.IsDigit(c) {
				if first_digit_found {
					second_digit = c
					second_digit_found = true
				} else {
					first_digit = c
					first_digit_found = true
				}
			}
			if len(word_cat) > 0 {
				for i, pn := range word_cat {
					word_idx[i] += 1
					if val, ok := spelling[pn]; ok {
						if c == val[word_idx[i]] {
							if word_idx[i] == len(val)-1 {
								if first_digit_found {
									second_digit = pn
									second_digit_found = true
								} else {
									first_digit = pn
									first_digit_found = true
								}
								word_cat[i] = 'X'
							}
						} else {
							word_cat[i] = 'X'
						}
					} else if rmap, ok := starters[pn]; ok {
						if rm, yes := rmap.(map[rune]rune); yes {
							if num, indeed := rm[c]; indeed {
								word_cat[i] = num
							} else {
								word_cat[i] = 'X'
							}
						}
					}
				}
			}
			if val, ok := starters[c]; ok {
				word_idx = append(word_idx, 0)
				if r, yes := val.(rune); yes {
					word_cat = append(word_cat, r)
				} else {
					word_cat = append(word_cat, c)
				}
			}
			// log.Println("Current 1st: ", first_digit, ", Current 2nd: ", second_digit)
		}
		if !second_digit_found {
			second_digit = first_digit
		}
		n, err := strconv.Atoi(string([]rune{first_digit, second_digit}))
		if err != nil {
			log.Println("Current line: ", scanner.Text())
			log.Println("Digits: ", string([]rune{first_digit, second_digit}))
			log.Fatal(err)
		}
		total += n
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return
}

func main() {
	p1 := partOne("input.txt")
	fmt.Println("Part One: ", p1)
	p2 := partTwo("input.txt")
	fmt.Println("Part Two: ", p2)
}
