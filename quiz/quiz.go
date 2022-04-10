package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"strings"
	"time"
)

func runQuiz(fileContents, timeLimit string) {
  input := make(chan string)
  csvReader := csv.NewReader(strings.NewReader(fileContents))

  duration, err := time.ParseDuration(timeLimit)
  if err != nil {
    panic("Could not parse the time limit specified.")
  }

  questions, correctAnswers := 0, 0
  for {
    line, err := csvReader.Read()
    if err == io.EOF {
      break
    }

    if err != nil {
      panic(err)
    }

    fmt.Print(line[0], ": ")
    go func () {
      var usersAnswer string
      fmt.Scan(&usersAnswer)
      input <- usersAnswer
    }()

    select {
    case answer := <-input:
      if line[1] == answer {
        correctAnswers++
      }
    case <-time.After(duration):
      fmt.Println()
    }

    questions++
  }

  close(input)

  fmt.Println("Total number of questions", questions)
  fmt.Println("Questions answered correctly", correctAnswers)
}

func main() {
  filePath := flag.String("file", "./problems.csv", "file path of problems")
  timeLimit := flag.String("tta", "30s", "time to answer each question")
  flag.Parse()

  fmt.Printf("Press any key to continue. You will have %s to answer each question.", *timeLimit)
  fmt.Scanln()

  file, err := ioutil.ReadFile(*filePath)
  if err != nil {
    panic("Input file does not exist or file path is wrong.")
  }

  contents := string([]byte(file))
  runQuiz(contents, *timeLimit)
}
