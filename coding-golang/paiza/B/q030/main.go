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

func move(c string, n *position, f *[][]string, w int, h int) {
	switch c {
	case "U":
		if n.y > 0 {
			n.y--
			iceMove(c, n, f, w, h)
			break
		}

		break
	case "D":
		if n.y < h-1 {
			n.y++
			iceMove(c, n, f, w, h)
			break
		}

		break
	case "L":
		if n.x > 0 {
			n.x--
			iceMove(c, n, f, w, h)
			break
		}

		break
	case "R":
		if n.x < w-1 {
			n.x++
			iceMove(c, n, f, w, h)
			break
		}

		break
	}
}

func iceMove(c string, n *position, f *[][]string, w int, h int) {
	if (*f)[n.y][n.x] == "#" {
		move(c, n, f, w, h)
	}
}

type position struct {
	x, y int
}

func main() {
	input := strings.Split(scanner(), " ")
	height, _ := strconv.Atoi(input[0])
	width, _ := strconv.Atoi(input[1])

	var field [][]string
	for i := 0; i < height; i++ {
		field = append(field, strings.Split(scanner(), ""))
	}

	input2 := strings.Split(scanner(), " ")
	x, _ := strconv.Atoi(input2[0])
	y, _ := strconv.Atoi(input2[1])

	var now position
	now.x = x - 1
	now.y = y - 1

	n, _ := strconv.Atoi(scanner())

	var command []string
	for i := 0; i < n; i++ {
		command = append(command, scanner())
	}

	for _, v := range command {
		move(v, &now, &field, width, height)
	}

	fmt.Println(now.x+1, now.y+1)
}
