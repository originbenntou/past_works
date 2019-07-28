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

func main() {
	i1 := strings.Split(scanner(), " ")
	m, _ := strconv.Atoi(i1[0])
	l, _ := strconv.Atoi(i1[1])

	var b int

	if m != 0 && l != 0 {
		var sum int
		for i := 0; i < l; i++ {
			i2 := strings.Split(scanner(), " ")
			for _, v := range i2 {
				s, _ := strconv.Atoi(v)
				sum += s
			}

			if sum > 0 {
				b += sum
			}

			sum = 0
		}
	}

	fmt.Println(b)
}
