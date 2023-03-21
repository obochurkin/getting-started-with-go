package main

import (
	"fmt"
	"strconv"
)

func main() {
  value := userInput();
	fmt.Printf("Truncated integer: %d \n", int32(*value))
}

func userInput() *float64 {
	var input string
	fmt.Println("Pls input a float number and hit Return")
	_, err := fmt.Scan(&input)
  parsedInput, e := strconv.ParseFloat(input, 32)
	if err != nil || e != nil {
    fmt.Println("Invalid input", err, e)
    userInput()
	}
  return &parsedInput
}
