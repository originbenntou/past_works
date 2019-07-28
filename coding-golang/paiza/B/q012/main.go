package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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

func even(v string) (c int) {
	converted, _ := strconv.Atoi(v)

	if c = converted * 2; c > 9 {
		c = 1 + (c % 10)
	}

	return
}

func main() {
	n, _ := strconv.Atoi(scanner())

	answer := make([]int, n)
	for i := 0; i < n; i++ {
		var sum = 0

		number := scanner()

		for k, v := range number {
			if k%2 == 0 {
				sum += even(string(v))
			} else {
				s := string(v)

				if s != "X" {
					converted, _ := strconv.Atoi(s)
					sum += converted
				}
			}
		}

		if a := sum % 10; a == 0 {
			answer[i] = 0
		} else {
			answer[i] = 10 - a
		}
	}

	for _, v := range answer {
		fmt.Println(v)
	}
}
