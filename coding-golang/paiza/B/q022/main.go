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
	input := strings.Split(scanner(), " ")
	m, _ := strconv.Atoi(input[0])
	n, _ := strconv.Atoi(input[1])
	k, _ := strconv.Atoi(input[2])

	voted := make([]int, m)
	voter := n
	for i := 0; i < k; i++ {
		order, _ := strconv.Atoi(scanner())
		getVote(&voted, order-1, getVoteNum(&voted, &voter, order-1))
	}

	var max int
	for _, v := range voted {
		if max < v {
			max = v
		}
	}

	for k, v := range voted {
		if max == v {
			fmt.Println(k + 1)
		}
	}

}

func getVoteNum(voted *[]int, voter *int, order int) (vote int) {
	for k, v := range *voted {
		if v > 0 && k != order {
			(*voted)[k]--
			vote++
		}
	}

	if *voter > 0 {
		*voter--
		vote++
	}

	return
}

func getVote(voted *[]int, order int, vote int) {
	(*voted)[order] += vote
	return
}
