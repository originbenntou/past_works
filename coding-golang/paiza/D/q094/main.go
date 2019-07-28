package main

import (
	"bufio"
	"fmt"
	"log"
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

func main() {
	var cat, dog int
	inputs := make([]string, 3)

	for i := 0; i < 3; i++ {
		inputs[i] = scanner()
	}

	for _, v := range inputs {
		if v == "cat" {
			cat++
		} else {
			dog++
		}
	}

	if cat > dog {
		fmt.Println("cat")
	} else {
		fmt.Println("dog")
	}
}
