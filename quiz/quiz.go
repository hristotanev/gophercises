package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"strings"
)

func main() {
  file, err := ioutil.ReadFile("problems.csv")
  if err != nil {
    panic("input file does not exist")
  }

  questions, correct := 0, 0
  contents := string([]byte(file))
  reader := csv.NewReader(strings.NewReader(contents))
  for {
    line, err := reader.Read()
    if err == io.EOF {
      break
    }

    if err != nil {
      panic(err)
    }

    fmt.Print("How much is ", line[0], ": ")
    var answer string
    fmt.Scan(&answer)

    if line[1] == answer {
      correct++
    }

    questions++
  }

  fmt.Println("Total number of questions", questions)
  fmt.Println("Questions answered correctly", correct)
}
