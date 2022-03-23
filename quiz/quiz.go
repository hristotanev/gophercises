package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"strings"
)

func part1(fileContents string) {
  questions, correctAnswers := 0, 0
  csvReader := csv.NewReader(strings.NewReader(fileContents))

  for {
    line, err := csvReader.Read()
    if err == io.EOF {
      break
    }

    if err != nil {
      panic(err)
    }

    fmt.Print("How much is ", line[0], ": ")
    var usersAnswer string
    fmt.Scan(&usersAnswer)

    if line[1] == usersAnswer {
      correctAnswers++
    }

    questions++
  }

  fmt.Println("Total number of questions", questions)
  fmt.Println("Questions answered correctly", correctAnswers)
}

func main() {
  filePath := flag.String("f", "./problems.csv", "file path of problems")
  flag.Parse()

  file, err := ioutil.ReadFile(*filePath)
  if err != nil {
    panic("input file does not exist")
  }

  contents := string([]byte(file))
  part1(contents)
}
