package models

type Question struct {
  Question string
  Answer string
}

func (question *Question) EvaluateAnswer(usersAnswer string) bool {
  if usersAnswer == question.Answer {
    return true
  }

  return false
}
