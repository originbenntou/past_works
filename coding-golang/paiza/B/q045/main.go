// FIXME: goroutinと必要な分だけ生成すること

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
	tasu, _ := strconv.Atoi(inputs[0])
	hiku, _ := strconv.Atoi(inputs[1])

	var sumFormula [100][100]interface{}
	var subFormula [100][100]interface{}

	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			if sum := i + j; sum < 100 {
				sumFormula[i][j] = i + j
			} else {
				sumFormula[i][j] = nil
			}
		}
	}

	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			if sub := i - j; sub > 0 {
				subFormula[i][j] = i - j
			} else {
				subFormula[i][j] = nil
			}
		}
	}

	count := 0
	for i := 0; i < 100; i++ {
		if count > tasu-1 {
			break
		}
		for j := 0; j < 100; j++ {
			if sumFormula[i][j] != nil {
				fmt.Println(strconv.Itoa(i) + " + " + strconv.Itoa(j) + " =")
				count++
			}

			if count > tasu-1 {
				break
			}
		}
	}

	count = 0
	for i := 0; i < 100; i++ {
		if count > hiku-1 {
			break
		}
		for j := 0; j < 100; j++ {
			if subFormula[i][j] != nil {
				fmt.Println(strconv.Itoa(i) + " - " + strconv.Itoa(j) + " =")
				count++
			}

			if count > hiku-1 {
				break
			}
		}
	}
}
