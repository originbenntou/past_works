/*
式の結果が合わないおわり
*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
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

const G = 9.8

func main() {
	input1 := strings.Split(scanner(), " ")

	o_y, _ := strconv.Atoi(input1[0])
	s, _ := strconv.Atoi(input1[1])
	rad, _ := strconv.Atoi(input1[2])

	input2 := strings.Split(scanner(), " ")

	x, _ := strconv.Atoi(input2[0])
	y, _ := strconv.Atoi(input2[1])
	diameter, _ := strconv.Atoi(input2[2])

	arrow_y := float64(o_y) + float64(x)*math.Tan(float64(rad)) -
		((G * math.Pow(float64(x), 2)) / (2 * math.Pow(float64(s), 2) * math.Pow(math.Cos(float64(rad)), 2)))

	fmt.Println(y, diameter, arrow_y)
}
