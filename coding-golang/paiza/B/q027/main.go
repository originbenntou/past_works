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
	H, _ := strconv.Atoi(inputs[0])
	W, _ := strconv.Atoi(inputs[1])
	N, _ := strconv.Atoi(inputs[2])

	player := make([]int, N)

	var tate [H][W]int
	for i := 0; i < H; i++ {
		yokoInput := strings.Split(scanner(), " ")

		for j := 0; j < W; j++ {
			tate[i][j], _ = strconv.Atoi(yokoInput[j])
		}
	}

	L, _ := strconv.Atoi(scanner())

	turn := 0
	for i := 0; i < L; i++ {
		mekuruInput := strings.Split(scanner(), " ")
		x1, _ := strconv.Atoi(mekuruInput[0])
		y1, _ := strconv.Atoi(mekuruInput[1])
		x2, _ := strconv.Atoi(mekuruInput[2])
		y2, _ := strconv.Atoi(mekuruInput[3])

		if tate[x1-1][y1-1] == tate[x2-1][y2-1] {
			player[turn] = player[turn] + 2
		} else {
			turn++

			if turn > N-1 {
				turn = 0
			}
		}
	}

	for _, v := range player {
		fmt.Println(v)
	}
}
