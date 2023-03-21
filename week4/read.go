package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

type Person struct {
	FirstName string `json:"firstName"`
  LastName string `json:"lastName"`
}

func main() {
  var fileName string
  fmt.Printf("Please enter created text file with a name pairs: ")
  _, err := fmt.Scan(&fileName)
  if err != nil {
    fmt.Println("Got an error", err)
	}
	content, err := ioutil.ReadFile(fileName)
  if err != nil {
    fmt.Println("got an error")
    panic(err)
  }

  names := strings.Split(string(content), "\n")

	for _, v := range names {
    firstLastNamePair := strings.Split(string(v), " ")
    p := Person{firstLastNamePair[0], firstLastNamePair[1]}
    j,_ := json.Marshal(p)
    fmt.Println(string(j))
	}
}
