package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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

var cards []int

func main() {
	inputs := strings.Split(scanner(), " ")

	for _, v := range inputs {
		number, _ := strconv.Atoi(v)
		cards = append(cards, number)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(cards)))

	fmt.Println(cards[0]*10 + cards[1]*10 + cards[2] + cards[3])
}
