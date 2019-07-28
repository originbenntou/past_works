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

type words []string

func (w words) Len() int {
	return len(w)
}

func (w words) Swap(i, j int) {
	w[i], w[j] = w[j], w[i]
}

func (w words) Less(i, j int) bool {
	return w[i] < w[j]
}

func main() {
	inputNum := strings.Split(scanner(), " ")
	wordsNum, _ := strconv.Atoi(inputNum[0])
	splitNum, _ := strconv.Atoi(inputNum[1])
	answerPage, _ := strconv.Atoi(inputNum[2])

	var answer [][]string
	if wordsNum%splitNum != 0 {
		answer = make([][]string, (wordsNum/splitNum)+1)
	} else {
		answer = make([][]string, wordsNum/splitNum)
	}

	var w words = strings.Split(scanner(), " ")
	sort.Sort(w)

	var pageNum, end = 0, 0

	for i := 0; i < len(w); i += splitNum {
		end = i + splitNum
		if len(w) < end {
			end = len(w)
		}
		answer[pageNum] = w[i:end]
		pageNum++
	}

	fmt.Println(strings.Join(answer[answerPage-1], "\n"))
}
