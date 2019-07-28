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
	_, _ = strconv.Atoi(scanner())
	numbers := strings.Split(scanner(), " ")

	var answer = 0
	for _, v := range numbers {
		me, _ := strconv.Atoi(v)

		answer += me
	}

	fmt.Println(answer)
}
