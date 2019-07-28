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

func calc(a [4]int, b [4]int, m [4]int, n int, w *[]int, x *[]int, y *[]int, z *[]int) []int {
	*w = append(*w, (a[0]*(*w)[n]+b[0])%m[0])
	*x = append(*x, (a[1]*(*x)[n]+b[1])%m[1])
	*y = append(*y, (a[2]*(*y)[n]+b[2])%m[2])
	*z = append(*z, (a[3]*(*z)[n]+b[3])%m[3])

	return []int{(*w)[n+1] % 10, (*x)[n+1] % 10, (*y)[n+1] % 10, (*z)[n+1] % 10}
}

func fortune(md []int, c []int) bool {
	for _, nn := range c {
		md = roop(nn, md)
	}

	return len(md) == 0
}

func roop(nn int, md []int) []int {
	for k, v := range md {
		if v == nn {
			return unset(md, k)
		}
	}

	return md
}

func unset(arr []int, n int) (hoge []int) {
	hoge = append(hoge, arr[:n]...)
	hoge = append(hoge, arr[n+1:]...)

	return
}

var w, x, y, z []int

func main() {
	input1 := strings.Split(scanner(), " ")

	aaa, _ := strconv.Atoi(input1[0])
	bbb, _ := strconv.Atoi(input1[1])

	var md []int
	if aaa < 10 {
		md = append(md, 0)
	} else {
		md = append(md, 1)
	}
	md = append(md, aaa%10)

	if bbb < 10 {
		md = append(md, 0)
	} else {
		md = append(md, 1)
	}
	md = append(md, bbb%10)

	input2 := strings.Split(scanner(), " ")

	a1, _ := strconv.Atoi(input2[0])
	a2, _ := strconv.Atoi(input2[1])
	a3, _ := strconv.Atoi(input2[2])
	a4, _ := strconv.Atoi(input2[3])

	var a = [4]int{a1, a2, a3, a4}

	input3 := strings.Split(scanner(), " ")

	b1, _ := strconv.Atoi(input3[0])
	b2, _ := strconv.Atoi(input3[1])
	b3, _ := strconv.Atoi(input3[2])
	b4, _ := strconv.Atoi(input3[3])

	var b = [4]int{b1, b2, b3, b4}

	input4 := strings.Split(scanner(), " ")

	m1, _ := strconv.Atoi(input4[0])
	m2, _ := strconv.Atoi(input4[1])
	m3, _ := strconv.Atoi(input4[2])
	m4, _ := strconv.Atoi(input4[3])

	var m = [4]int{m1, m2, m3, m4}

	var n = 0
	w = append(w, 0)
	x = append(x, 0)
	y = append(y, 0)
	z = append(z, 0)

	var isLucky = false
	for !isLucky {
		cards := calc(a, b, m, n, &w, &x, &y, &z)
		isLucky = fortune(md, cards)
		if !isLucky {
			n++
		} else {
			fmt.Println(n + 1)
		}
	}
}
