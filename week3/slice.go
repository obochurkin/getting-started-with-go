package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	sli := []int{3, 34, 11}
	fmt.Println("Pls input integer and hit Return")
	scanner := bufio.NewScanner(os.Stdin)

	for true {
		scanner.Scan()
		input := scanner.Text()

		if len([]rune(input)) < 1 {
			fmt.Println("Got empty string")
      continue
		}
		number, err := strconv.ParseInt(strings.TrimSpace(input), 10, 64)

		if err != nil {
			fmt.Println("Got an error", err)
			return
		}

		sli = append(sli, int(number))

		sort.Ints(sli)
		fmt.Println(sli)
	}
}
