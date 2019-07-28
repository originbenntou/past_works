package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"
)

var f = os.Stdin
var stdin = bufio.NewScanner(os.Stdin)

func scanner() (s string) {
	//var f = os.Stdin
	//var stdin = bufio.NewScanner(f)

	fmt.Println("===Before Scan========")

	// ファイル現在行の取得
	fmt.Println("File Line:", liner(f))
	// 構造体の中身
	analyzer(stdin)

	if stdin.Scan() {
		fmt.Println("===After Scan========")

		// ファイル現在行の取得
		fmt.Println("File Line:", liner(f))
		// 構造体の中身
		analyzer(stdin)

		s = strings.TrimSpace(stdin.Text())
	} else {
		log.Fatalln(stdin.Err())
	}

	return
}

func analyzer(stdin *bufio.Scanner) {
	v := reflect.Indirect(reflect.ValueOf(stdin))
	t := v.Type()

	fmt.Println("***Struct***")
	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).Name != "buf" {
			fmt.Println(t.Field(i).Name, v.Field(i))
		}
	}

	return
}

func liner(f *os.File) (line int64) {
	line, _ = f.Seek(0, 1)
	return
}

func main() {
	for i := 0; i < 3; i++ {
		scanner()
	}
}
