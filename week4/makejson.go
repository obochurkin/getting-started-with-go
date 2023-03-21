package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strings"
)

type Person struct {
	Name string `json:"name"`
  Address string `json:"address"`
}

func main() {
	var p Person
	scanner := bufio.NewScanner(os.Stdin)

  names := getKeys(&p)

  inputValues := make(map[string]string)

  for _, v := range names {
    inputValues[v] = requestInputData(scanner, v)
  }

  jsonMapData, _ := json.Marshal(inputValues)
  json.Unmarshal(jsonMapData, &p) // converting map to object
  jsonStructData, _ := json.Marshal(p)
  fmt.Println(string(jsonStructData))
}

func requestInputData(scanner *bufio.Scanner, fieldName string) string {
  fmt.Printf("Pls input value for the %s and hit Return \n", fieldName)
  scanner.Scan()
	input := scanner.Text()

  if len([]rune(input)) < 1 {
    fmt.Println("Got empty string")
    requestInputData(scanner, fieldName)
  }

  return strings.TrimSpace(input)
}

func getKeys[T comparable](s *T) []string {
  var keys []string
  v := reflect.ValueOf(s).Elem()
  for i:=0; i < v.NumField(); i++ {
      keys = append(keys, v.Type().Field(i).Name)
  }

  return keys
}
