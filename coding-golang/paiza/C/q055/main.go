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
	N, _ := strconv.Atoi(scanner())
	G := scanner()

	var logs = []string{"a", "b"}
	fmt.Println(len(logs))
	for i := 0; i < N; i++ {
		logs[i] = scanner()
	}

	var count int
	for _, v := range logs {
		if strings.Contains(v, G) {
			fmt.Println(v)
			count++
		}
	}

	if count == 0 {
		fmt.Println("None")
	}
}
