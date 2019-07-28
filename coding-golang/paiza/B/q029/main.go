/*
持ちたいデータは構造体で定義して、その集合体をスライスで定義する感じが理解できた
このように考えればソートは、構造体のどの値でソートしたいかだけ考えればよいので難しくない
あと、GoにRoundはない
*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
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

func calc(x1 float64, y1 float64, x2 float64, y2 float64) (distance float64) {
	distance = math.Sqrt(math.Pow(x1-x2, 2.0) + math.Pow(y1-y2, 2.0))

	return
}

type chika struct {
	d float64
	p float64
}

type chikaList []chika

func (c chikaList) Len() int {
	return len(c)
}

func (c chikaList) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c chikaList) Less(i, j int) bool {
	return c[i].d < c[j].d
}

func round(f float64) float64 {
	return math.Floor(f + .5)
}

func main() {
	a := strings.Split(scanner(), " ")
	ax, _ := strconv.ParseFloat(a[0], 64)
	ay, _ := strconv.ParseFloat(a[1], 64)

	k, _ := strconv.Atoi(scanner())
	N, _ := strconv.Atoi(scanner())

	list := make(chikaList, N)
	for i := 0; i < N; i++ {
		inputs := strings.Split(scanner(), " ")
		tx, _ := strconv.ParseFloat(inputs[0], 64)
		ty, _ := strconv.ParseFloat(inputs[1], 64)
		p, _ := strconv.ParseFloat(inputs[2], 64)

		list[i].d = calc(ax, ay, tx, ty)
		list[i].p = p
	}

	sort.Sort(list)

	var sum float64
	for i := 0; i < k; i++ {
		sum += list[i].p
	}

	fmt.Println(round(sum / float64(k)))
}
