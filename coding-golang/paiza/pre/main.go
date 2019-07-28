package pre

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

func checkYear(y int) (b bool) {
	if (y % 4) == 0 {
		if (y % 400) == 0 {
			b = true
			return
		}

		if (y % 100) == 0 {
			b = false
			return
		}

		b = true
	}

	return
}

func main() {
	lines, err := strconv.Atoi(scanner())
	if err != nil {
		log.Fatalln(err)
	}

	answer := make([]string, lines)

	for i := 0; i < lines; i++ {
		yearString := scanner()
		yearInt, err := strconv.Atoi(yearString)
		if err != nil {
			log.Fatalln(err)
		}

		if checkYear(yearInt) {
			answer[i] = yearString + " is a leap year"
		} else {
			answer[i] = yearString + " is not a leap year"
		}
	}

	fmt.Println(strings.Join(answer, "\n"))
}
