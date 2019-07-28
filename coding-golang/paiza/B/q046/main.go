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

func main() {
	_, _ = strconv.Atoi(scanner())

	input1 := strings.Split(scanner(), " ")
	input2 := strings.Split(scanner(), " ")

	// 現在地
	var n1, d1 int
	if n1, _ = strconv.Atoi(input1[0]); input1[1] == "W" || input1[1] == "S" {
		n1 = -n1
	}

	if input1[1] == "N" || input1[1] == "S" {
		d1 = 0
	} else {
		d1 = 1
	}

	// 目的地
	var n2, d2 int
	if n2, _ = strconv.Atoi(input1[0]); input2[1] == "W" || input2[1] == "S" {
		n2 = -n2
	}

	if input2[1] == "N" || input2[1] == "S" {
		d2 = 0
	} else {
		d2 = 1
	}

	var distance float64
	if d1 == d2 {
		// 直線
		distance = 100 * math.Abs(float64(n1)-float64(n2))
	} else {
		// 円周上
		var higher, lower, depthDiff float64
		higher = math.Abs(float64(n1))
		lower = math.Abs(float64(n2))

		if depthDiff := higher - lower; depthDiff < 0 {
			higher = lower
		}

		distance = float64(higher)*100*math.Pi*0.5 + float64(depthDiff)*100
	}

	fmt.Println(distance)
}
