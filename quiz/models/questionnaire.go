package models

import "fmt"

type Question struct {
  Question string
  Answer string
}

type Questionnaire struct {
  Questions []Question
  CorrectlyAnsweredQuestions int
}

func (questionnaire *Questionnaire) EvaluateAnswer(usersAnswer string, question Question) {
  if usersAnswer == question.Answer {
    questionnaire.CorrectlyAnsweredQuestions++
  }
}

func (questionnaire *Questionnaire) PrintSummary() {
  fmt.Println("Total number of questions", len(questionnaire.Questions))
  fmt.Println("Questions answered correctly", questionnaire.CorrectlyAnsweredQuestions)
}
