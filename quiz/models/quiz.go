package models

import (
	"fmt"
	"math/rand"
	"time"
)

type Quiz struct {
  Questions []Question
  CorrectlyAnsweredQuestions int
}

func (quiz *Quiz) AddQuestionAnswer(isQuestionAnswerCorrect bool) {
  if isQuestionAnswerCorrect {
    quiz.CorrectlyAnsweredQuestions++
  }
}

func (quiz *Quiz) ShuffleQuestions() {
  rand.Seed(time.Now().UnixNano())
  rand.Shuffle(len(quiz.Questions), func (i, j int) {
    quiz.Questions[i], quiz.Questions[j] = quiz.Questions[j], quiz.Questions[i]
  })
}

func (quiz *Quiz) PrintSummary() {
  fmt.Println("Total number of questions", len(quiz.Questions))
  fmt.Println("Questions answered correctly", quiz.CorrectlyAnsweredQuestions)
}
