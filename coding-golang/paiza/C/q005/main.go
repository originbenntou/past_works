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

	a := make([]string, N)
	for i := 0; i < N; i++ {
		input := strings.Split(scanner(), ".")
		if len(input) > 4 {
			a[i] = "False"
		} else {
			for _, v := range input {
				o, e := strconv.Atoi(v)
				if e != nil {
					a[i] = "False"
					break
				}

				if o > 255 || o < 0 {
					a[i] = "False"
					break
				}

				a[i] = "True"
			}
		}

	}

	for _, v := range a {
		fmt.Println(v)
	}
}
