package models

type Question struct {
  Question string
  Answer string
}

func (question *Question) EvaluateAnswer(usersAnswer string) bool {
  return usersAnswer == question.Answer
}
