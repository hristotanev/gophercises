package models

import (
	"fmt"
	"math/rand"
	"time"
)

type Question struct {
  Question string
  Answer string
}

type Quiz struct {
  Questions []Question
  CorrectlyAnsweredQuestions int
}

func (quiz *Quiz) ShuffleQuestions() {
  rand.Seed(time.Now().UnixNano())
  rand.Shuffle(len(quiz.Questions), func (i, j int) {
    quiz.Questions[i], quiz.Questions[j] = quiz.Questions[j], quiz.Questions[i]
  })
}

func (quiz *Quiz) EvaluateAnswer(usersAnswer string, question Question) {
  if usersAnswer == question.Answer {
    quiz.CorrectlyAnsweredQuestions++
  }
}

func (quiz *Quiz) PrintSummary() {
  fmt.Println("Total number of questions", len(quiz.Questions))
  fmt.Println("Questions answered correctly", quiz.CorrectlyAnsweredQuestions)
}
