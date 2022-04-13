package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	. "quiz/models"
	"strings"
	"time"
)

var inputChannel chan string

func getQuestionsFromCSV(fileContents string) []Question {
  csvReader := csv.NewReader(strings.NewReader(fileContents))

  var questions []Question
  for {
    line, err := csvReader.Read()
    if err == io.EOF {
      break
    }

    if err != nil {
      panic(err)
    }

    questions = append(questions, Question{
      Question: line[0],
      Answer: line[1],
    })
  }

  return questions
}

func startQuiz(fileContents string, duration time.Duration, shouldShuffleQuestions bool) {
  quiz := &Quiz{
    Questions: getQuestionsFromCSV(fileContents),
    CorrectlyAnsweredQuestions: 0,
  }

  if shouldShuffleQuestions {
    quiz.ShuffleQuestions()
  }

  for _, question := range quiz.Questions {
    fmt.Print(question.Question, ": ")
    go func () {
      var answer string
      fmt.Scan(&answer)
      inputChannel <- strings.ToLower(strings.TrimSpace(answer))
    }()

    select {
    case answer := <-inputChannel:
      quiz.EvaluateAnswer(answer, question)
    case <-time.After(duration):
      fmt.Println()
    }
  }

  quiz.PrintSummary()
}

func main() {
  filePath := flag.String("file", "./problems.csv", "file path of problems")
  timeLimit := flag.String("tta", "30s", "time to answer each question")
  shouldShuffleQuestions := flag.Bool("shfl", false, "enable question shuffle")
  flag.Parse()

  file, err := ioutil.ReadFile(*filePath)
  if err != nil {
    panic("Input file does not exist or file path is wrong.")
  }

  duration, err := time.ParseDuration(*timeLimit)
  if err != nil {
    panic("Could not parse the time limit specified.")
  }

  fmt.Printf("Press any key to continue. You will have %s to answer each question.", *timeLimit)
  fmt.Scanln()

  inputChannel = make(chan string)

  contents := string([]byte(file))
  startQuiz(contents, duration, *shouldShuffleQuestions)

  close(inputChannel)
}
