package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

var stdin = bufio.NewScanner(os.Stdin)

func scanner() (s string) {
	if stdin.Scan() {
		s = strings.TrimSpace(stdin.Text())
	} else {
		log.Fatalln(stdin.Err())
	}

	return
}

type typo struct {
	hand int
	key  string
	row  int
	col  int
}

var keyboard [3][]string
var history []typo

var now typo

func main() {
	keyboard[0] = strings.Split("qwertyuiop", "")
	keyboard[1] = strings.Split("asdfghjkl ", "")
	keyboard[2] = strings.Split("zxcvbnm   ", "")

	input := strings.Split(scanner(), "")

	for _, c := range input {
		now.key = string(c)
		for r, line := range keyboard {
			for c, k := range line {
				if now.key == k {
					now.row = r
					now.col = c
					if c < 5 {
						now.hand = 0
					} else {
						now.hand = 1
					}
				}
			}
		}
		history = append(history, now)
	}

	var count int
	for i := 0; i < len(history); i++ {
		if i == len(history)-1 {
			break
		}

		if math.Abs(float64(history[i].row)-float64(history[i+1].row)) < float64(2) && history[i].col == history[i+1].col ||
			math.Abs(float64(history[i].col)-float64(history[i+1].col)) < float64(2) && history[i].row == history[i+1].row {
			if history[i].hand != history[i+1].hand {
				count++
				history[i+1].hand = history[i].hand
			}
		}
	}

	fmt.Println(count)
}
