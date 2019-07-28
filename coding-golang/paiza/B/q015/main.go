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

var numbers = [10]string{
	"1111110",
	"0110000",
	"1101101",
	"1111001",
	"0110011",
	"1011011",
	"1011111",
	"1110010",
	"1111111",
	"1111011",
}

func isNumber(i string) (b bool) {
	for _, v := range numbers {
		if b = i == v; b {
			return b
		}
	}

	return b
}

func taisyo(i string) (s string) {
	return string(i[0]) + string(i[5]) + string(i[4]) + string(i[3]) + string(i[2]) + string(i[1]) + string(i[6])
}

func kaiten(i string) (s string) {
	return string(i[3]) + string(i[4]) + string(i[5]) + string(i[0]) + string(i[1]) + string(i[2]) + string(i[6])
}

func main() {
	input1 := strings.Replace(scanner(), " ", "", -1)
	input2 := strings.Replace(scanner(), " ", "", -1)

	if r := isNumber(input1) && isNumber(input2); r {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
	if r := isNumber(taisyo(input1)) && isNumber(taisyo(input2)); r {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
	if r := isNumber(kaiten(input1)) && isNumber(kaiten(input2)); r {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
