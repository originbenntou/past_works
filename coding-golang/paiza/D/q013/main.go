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
	inputs := strings.Split(scanner(), " ")
	m, _ := strconv.Atoi(inputs[0])
	n, _ := strconv.Atoi(inputs[1])

	sho := m / n
	amari := m % n

	fmt.Println(sho, amari)
}
