package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	var input string
  var re = regexp.MustCompile(`i+a+n`)
  fmt.Println("Welcome /n This program will detect letters i, a, n and letting user know if those letter were found in the provided string")
	fmt.Println("Pls input string and hit Return")
	_, err := fmt.Scan(&input)
	if err != nil {
    fmt.Println("Got an error", err)
	}
  if re.MatchString(strings.TrimSpace(strings.ToLower(input))) {
    fmt.Println("Found!")
    return
  }
  fmt.Println("Not Found!")
}
